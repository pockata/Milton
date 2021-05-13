import { addDecorator } from '@storybook/react';
import { initializeWorker, mswDecorator } from 'msw-storybook-addon';

// Sometimes SWR returns stale data. Use swr.cache.clear()
// https://github.com/vercel/swr/issues/27#issuecomment-743196474
// https://storybook.js.org/docs/react/addons/addons-api#apioneventname-fn
// https://mswjs.io/docs/faq#swr
// addons.register('lek', (api) => {
// 	console.table(api)
// 	api.on('keydown', (eventData) => {
// 		console.log(eventData);
// 	});
// })

initializeWorker();
addDecorator(mswDecorator);

export const parameters = {
	actions: { argTypesRegex: "^on[A-Z].*" },
	controls: {
		matchers: {
			color: /(background|color)$/i,
			date: /Date$/,
		},
	},
}

