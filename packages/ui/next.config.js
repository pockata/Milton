var DirectoryNamedWebpackPlugin = require("directory-named-webpack-plugin");

module.exports = {
	/**
	 * @param {import('webpack').Configuration} config
	 */
	webpack: (config) => {
		if (!config.resolve) {
			config.resolve = {plugins: []};
		}
		// @ts-ignore
		config.resolve.plugins.push(new DirectoryNamedWebpackPlugin());

		// Important: return the modified config
		return config;
	},
};

