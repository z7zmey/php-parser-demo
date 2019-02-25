var editor = ace.edit("editor", {
    mode: "ace/mode/php",
    selectionStyle: "text",
    autoScrollEditorIntoView: true,
    copyWithEmptySelection: true,
    theme: "ace/theme/chrome",
    showGutter: false,
})

var output = ace.edit("output", {
    mode: "ace/mode/golang",
    selectionStyle: "text",
    autoScrollEditorIntoView: true,
    copyWithEmptySelection: true,
    theme: "ace/theme/chrome",
    showGutter: false,
    readOnly: true,
})

var wto;
editor.on("change", function(e) {
    clearTimeout(wto);
    wto = setTimeout(function() {
        requestToParser()
    }, 300);
});

$('input.switch-imput').on("change", function(e) {
    requestToParser()
});

function requestToParser(src) {
    console.log("test")
    var src = editor.getValue()
    var php5 = $('#switch-to-php5').is(':checked')
    var freefloating = $('#switch-to-show-free-floating').is(':checked')

    $.ajax({
        url: "/parse",
        data: {
            script: src,
            php5: php5,
            free_floating: freefloating,
        },
        type: 'POST',
        success: function(result) {
            output.setValue(result);
            output.clearSelection()
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