export async function sendRequest(type, url, data) {
  try {
    const res = await fetch(url, {
      method: type,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    if (res.ok) {
      return await res.json();
    } else {
      const { message } = await res.json();
      new Error(message);
    }
  } catch(error) {
    console.error(error);
    return {
      success: false,
      result: error
    }
  }
}
