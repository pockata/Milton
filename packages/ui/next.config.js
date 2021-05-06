var DirectoryNamedWebpackPlugin = require("directory-named-webpack-plugin");

module.exports = {
	webpack: (config) => {
		if (!config.resolve) {
			config.resolve = {plugins: []};
		}
		config.resolve.plugins.push(new DirectoryNamedWebpackPlugin());

		// Important: return the modified config
		return config;
	},
};

