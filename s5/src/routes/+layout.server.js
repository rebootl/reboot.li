
/** @type {import('./$types').LayoutServerLoad} */

export const load = (async ({ locals }) => {

  console.log('locals', locals);
  if (locals.user) {
    return {
      loggedIn: true,
      username: locals.user.name,
    }
  } else {
    return {
      loggedIn: false,
      username: null,
    }
  }
});
