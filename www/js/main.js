var editor = ace.edit("editor", {
    mode: "ace/mode/php",
    selectionStyle: "text",
    autoScrollEditorIntoView: true,
    copyWithEmptySelection: true,
    theme: "ace/theme/chrome",
})

var wto;
editor.on("change", function(e) {
    clearTimeout(wto);
    wto = setTimeout(function() {
        var src = editor.getValue()
        requestToParser(src)
    }, 300);
});

function requestToParser(src) {
    $.ajax({
        url: "/parse",
        data: {
            script: src
        },
        type: 'POST',
        success: function(result) {
            $("#output").text(result);
        }
    });
}

editor.insert(`<?php

namespace Foo;

abstract class Bar extends Baz
{
	public function greet()
	{
		echo "Hello World";
	}
}`);