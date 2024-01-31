export async function sendJSONRequest(type, url, data) {
  try {
    const res = await fetch(url, {
      method: type,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      credentials: 'include',
      body: JSON.stringify(data)
    });
    if (res.ok) {
      return {
        success: true,
        result: await res.json()
      }
    } else {
      throw(new Error(`${res.status} ${res.statusText}`));
    }
  } catch(error) {
    console.error(error);
    return {
      success: false,
      result: error
    }
  }
}

export async function sendTokenRequest(type, url, data, token) {
  try {
    const res = await fetch(url, {
      method: type,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token,
      },
      body: JSON.stringify(data)
    });
    if (res.ok) {
      return await res.json();
    } else {
      throw(new Error(`${res.status} ${res.statusText}`));
    }
  } catch(error) {
    console.error(error);
    return {
      success: false,
      result: error
    }
  }
}

export async function getToken() {
  try {
    const res = await fetch('/getToken');
    if (!res.ok) {
      return false;
    }
    const j = await res.json();
    return j.token;
  } catch (error) {
    console.error(error);
    return false;
  }
}
