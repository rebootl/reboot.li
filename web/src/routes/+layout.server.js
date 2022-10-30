export async function load({ locals }) {
  /*return {
    user: await db.getUser(request.headers.get('cookie'))
  };*/

	//console.log(locals)

  if (locals.loggedIn) {
    return {
      loggedIn: true,
      user: locals.user,
      admin: locals.admin,
    }
  } else {
    return {
      loggedIn: false,
    }
  }
}