import { rest } from 'msw';
import React from 'react';
import ListUnits from './ListUnits.jsx';

/** @typedef {import('@storybook/react').Meta} Meta  */
/** @typedef {import('@storybook/react').Story} Story  */

// This default export determines where your story goes in the story list
/** @type {Meta} */
export default {
	title: 'ListUnits',
	component: ListUnits,
};

// We create a “template” of how args map to rendering
const Template = () => <ListUnits />;

/** @type Story */
export const Default = Template.bind({});

Default.parameters = {
	msw: [
		rest.get('/api/get-all-units', (_req, res, ctx) => {
			return res(ctx.json({
				"error": false,
				"data": [
					{ "ID": 1, "Name": "Na terasata", "MDNS": "Milton-1234", "Pots": null },
					{ "ID": 2, "Name": "Na terasata", "MDNS": "Milton-2", "Pots": null }
				]
			}));
		}),
	],
};

/** @type Story */
export const NoData = Template.bind({});

NoData.parameters = {
	msw: [
		rest.get('/api/get-all-units', (_req, res, ctx) => {
			return res(ctx.json({
				"error": false,
				"data": []
			}));
		}),
	],
};

/** @type Story */
export const Error = Template.bind({});

Error.parameters = {
	msw: [
		rest.get('/api/get-all-units', (_req, res, ctx) => {
			return res(ctx.json({
				"error": true,
				"data": []
			}));
		}),
	],
};

