import Compressor from 'compressorjs';
import { session } from '$app/stores';

export const compressImage = (file, maxWidth=1920, maxHeight=1920) => {
  return new Promise((res, rej) => {
    new Compressor(file, {
      maxWidth: maxWidth,
      maxHeight: maxHeight,
      success(result) {
        res(result);
      },
      error(err) {
        rej(err);
      }
    });
  });
};

// base64 encode data
export const encodeData = (file) => {
  return new Promise((res, rej) => {
    const reader = new FileReader();
    reader.onloadend = () => {
      res(reader.result);
    }
    reader.readAsDataURL(file);
  });
};

// upload w/ progress
export async function *uploadMultiImagesGenerator(images) {
  const files = await Promise.all(images.map(async (i) => {
    const blob = await compressImage(i.file, 1920, 1920);
    return new File([blob], i.filename);
  }));
  for await (const r of uploadMultiFilesGenerator(`http://localhost:3005/api/uploadImages`, files)) {
    yield r;
  }
}

export async function* uploadMultiFilesGenerator(apiUrl, files) {
  const formData = new FormData();
  for (const f of files) {
    formData.append('filedata', f);
  }
  const xhr = new XMLHttpRequest();
  let progress = 0.;
  let done = false;
  let result = {};
  let res = () => {};
  let p = new Promise((r) => res = r);
  const update = () => {
    res();
    p = new Promise((r) => res = r);
  };
  xhr.upload.addEventListener('progress', (e) => {
    progress = (e.loaded / e.total) * 100;
    update();
  });
  xhr.addEventListener('load', (e) => {
    result = xhr.response;
    done = true;
    update();
  });
  xhr.addEventListener('error', (e) => {
    console.log("Error during xhr transfer...", xhr.response);
    done = true;
    update();
  });
  xhr.addEventListener('abort', (e) => {
    console.log("Upload aborted...", xhr.response);
    done = true;
    update();
  });
  xhr.responseType = 'json';
  xhr.open('post', apiUrl);
  xhr.withCredentials = true;
  //xhr.setRequestHeader('Authorization', getAuthHeader()['Authorization']);
  xhr.send(formData);
  while(!done) {
    await p;
    yield {
      progress: progress,
      result: result,
      request: xhr
    };
  }
}
