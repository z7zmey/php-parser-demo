var editor = ace.edit("editor", {
    mode: "ace/mode/php",
    selectionStyle: "text",
    autoScrollEditorIntoView: true,
    copyWithEmptySelection: true,
    theme: "ace/theme/chrome",
    showGutter: false,
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
    var src = editor.getValue()
    var php5 = $('#switch-to-php5').is(':checked')
    var positions = $('#switch-to-show-pos').is(':checked')
    var comments = $('#switch-to-show-comments').is(':checked')

    $.ajax({
        url: "/parse",
        data: {
            script: src,
            php5: php5,
            positions: positions,
            comments: comments,
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