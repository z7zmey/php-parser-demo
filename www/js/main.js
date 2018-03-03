var editor = ace.edit("editor", {
    mode: "ace/mode/php",
    selectionStyle: "text",
    autoScrollEditorIntoView: true,
    copyWithEmptySelection: true,
    theme: "ace/theme/chrome",
})

editor.insert(`<?php

namespace Foo;

abstract class Bar extends Baz
{
	public function greet()
	{
		echo "Hello World";
	}
}`);