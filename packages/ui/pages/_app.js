import React from 'react';
import PropTypes from 'prop-types';

import "normalize.css";
import '../styles/globals.scss';

import { SWRConfig } from 'swr';
import config from '../config';

/**
 * @param {object} props
 * @param {React.ComponentType} props.Component
 * @param {object} props.pageProps
 *
 * @return {React.ReactElement}
 */
function MyApp({ Component, pageProps }) {
	/**
	 * @param {string} resource
	 * @param {RequestInit} [init]
	 **/
	let fetcher = (resource, init) => fetch(config.apiURL + resource, init).then(res => res.json());
	return (
		<SWRConfig value={{fetcher}} >
			<Component {...pageProps} />
		</SWRConfig>
	);
}

MyApp.propTypes = {
	Component: PropTypes.instanceOf(React.Component),
	pageProps: PropTypes.object,
};

export default MyApp;

