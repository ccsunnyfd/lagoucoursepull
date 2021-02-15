package template

// LessonTextTemplate LessonTextTemplate
var LessonTextTemplate = `
<html>
	<head>
		<meta charset="utf-8">
		<title>$LESSON_TITLE</title>
		<style>
			body {
				line-height: 1.8;
				font-family: -apple-system,BlinkMacSystemFont,Helvetica Neue,Helvetica,Segoe UI,Arial,Roboto,PingFang SC,miui,Hiragino Sans GB,Microsoft Yahei,sans-serif;
				font-size: 20px;
				height: 100%;
				width: 80%;
				background-color: #DCE2F1;
				border-radius: 4px;	
				margin: auto auto;
			}
			img {
				width: 600px;
				margin: 0 200px;
			}
			.main-title {
				font-weight: 700;
				font-size: 24px;
				color: #00b38a;
				margin-bottom: 8px;
			}
			.rich-text-wrap {
				font-size: 19px;
				margin: 50px auto;
				padding-bottom: 50px;
			}
			.rich-text-wrap pre {
				margin: .5em 0;
				padding: .4em .6em;
				border-radius: 8px;
				background: #000033;
				width: 716px;
				overflow-x: auto;
				box-sizing: border-box;
				color: #f6f8fa;
			}
			.course-code-area {
				position: relative;
			}
			code {
				font-family: monospace;
			}
			pre {
				display: block;
				font-family: monospace;
				white-space: pre;
			}
			.rich-text-wrap .hljs-bullet, .rich-text-wrap .hljs-string, .rich-text-wrap .hljs-symbol {
				color: #60ac39;
			}
			.rich-text-wrap .hljs-section, .rich-text-wrap .hljs-title {
				color: #6684e1;
			}
			.rich-text-wrap .hljs-keyword, .rich-text-wrap .hljs-selector-tag {
				color: #b854d4;
			}
			.rich-text-wrap .hljs-built_in, .rich-text-wrap .hljs-builtin-name, .rich-text-wrap .hljs-literal, .rich-text-wrap .hljs-meta, .rich-text-wrap .hljs-number, .rich-text-wrap .hljs-params, .rich-text-wrap .hljs-type {
				color: #b65611;
			}
		</style>
	</head>
	<body>
		<div data-v-39db28e6 class="rich-text-wrap">
		<div class="main-title">
			$LESSON_TITLE
		</div>
			$SOURCE_HTML_TEXT
	</div>
	</body>
</html>
`
