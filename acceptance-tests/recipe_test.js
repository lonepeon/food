const { above, attach, openBrowser, closeBrowser, goto, click, fileField, into, screenshot, setConfig, text, textBox, timeField, waitFor, write } = require("taiko");
const assert = require("assert").strict;
const settings = require("./settings").load();
const helpers = require("./helpers");
const path = require("path");

const browserArgs = {args:['--no-sandbox', '--disable-setuid-sandbox']};

describe("say hello", () => {
	beforeEach(async() => {
		await openBrowser(browserArgs);
		await goto(settings.url);
	});

	afterEach(async() => {
		await closeBrowser();
	});

	describe("say hello", () => {
		helpers.uiTest("works", async () => {
			await assert.ok(await text("Hello World").exists(0,0));
		});
	});
});
