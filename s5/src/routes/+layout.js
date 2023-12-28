
/** @type {import('./$types').LayoutLoad} */
export async function load({ data }) {

  const clientData = data;

  // console.log('clientData', clientData);

  return {
    clientData,
  };
}
