// The file contents for the current environment will overwrite these during build.
// The build system defaults to the dev environment which uses `environment.ts`, but if you do
// `ng build --env=prod` then `environment.prod.ts` will be used instead.
// The list of which env maps to which file can be found in `.angular-cli.json`.

export const environment = {
    production: false,
    api: {
        url: 'http://localhost:8096',
        client_id: '23047803-2a7e-46c1-adf8-76cd42f9518d'
    },
};