import { cache } from 'swr';
import { addons } from '@storybook/addons';
import { SELECT_STORY, CURRENT_STORY_WAS_SET } from '@storybook/core-events';

// Sometimes SWR returns stale data. Use swr.cache.clear()
// https://github.com/vercel/swr/issues/27#issuecomment-743196474
// https://storybook.js.org/docs/react/addons/addons-api#apioneventname-fn
// https://github.com/storybookjs/storybook/blob/master/lib/core-events/src/index.ts
// https://mswjs.io/docs/faq#swr
addons.register('lek', (api) => {
	console.table(api);
	api.on(SELECT_STORY, () => {
		cache.clear();
	});
	api.on(CURRENT_STORY_WAS_SET, () => {
		cache.clear();
	});
})

