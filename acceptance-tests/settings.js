exports.load = () => {
	const variables = {
		url: "ACCEPTANCE_TEST_SERVER_URL",
	};

	return Object.entries(variables).reduce((o, [attrName, envVarName]) => {
		const envVarValue = process.env[envVarName];
		if (!envVarValue) {
			throw `${envVarName} environment variable is not set`
		}

		o[attrName] = envVarValue;

		return o;
	}, {});
}
