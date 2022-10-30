export async function load({ data }) {
  /*return {
    user: await db.getUser(request.headers.get('cookie'))
  };*/

	//console.log(data)

  return {
    ...data
  }
}