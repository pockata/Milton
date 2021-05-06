module.exports = {
    "env": {
        "browser": true,
        "es2021": true,
        "node": true
    },
    "extends": [
		// See https://github.com/gajus/eslint-plugin-jsdoc
        "eslint:recommended",
        "plugin:react/recommended",
        "plugin:@next/next/recommended"
    ],
    "parserOptions": {
        "ecmaFeatures": {
            "jsx": true
        },
        "ecmaVersion": 12,
        "sourceType": "module"
    },
    "plugins": [
        "react"
    ],
    "rules": {
        "react/prop-types": [2, { ignore: ['children'] }],
        "react/react-in-jsx-scope": "off",
		"semi": "error",
    }
};
