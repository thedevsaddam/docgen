$(document).ready(function(){
        // Add minus icon for collapse element which is open by default
        $(".collapse.in").each(function(){
        	$(this).siblings(".panel-heading").find(".glyphicon").addClass("glyphicon-minus").removeClass("glyphicon-plus");
        });
        
        // Toggle plus minus icon on show hide of collapse element
        $(".collapse").on('show.bs.collapse', function(){
        	$(this).parent().find(".glyphicon").removeClass("glyphicon-plus").addClass("glyphicon-minus");
        }).on('hide.bs.collapse', function(){
        	$(this).parent().find(".glyphicon").removeClass("glyphicon-minus").addClass("glyphicon-plus");
        });

        
        $('.resp-prettyprint').each(function() {
        	var ctx = $(this);
        	var html =  ctx.html();
		    ctx.html("");
		    
		    html = html.replaceAll('\\n', '');
		    html = html.replaceAll('\\t', '');
		    html = html.replaceAll('\\', '');
		    html = html.trim();

		    if (html.charAt(0) === '"'){
		    	html = html.substr(1);
		    	html = html.slice(0, -1);
            }
            if (IsJsonString(html)){
                var obj = JSON.parse(html);
                var formattedJson = JSON.stringify(obj, null, 4);
                ctx.html("<pre>" + syntaxHighlight(formattedJson) + "</pre>");
            } else {
                ctx.html("<pre>" + escapeHtml(html) + "</pre>");
            }
		});

        $(".resp-selector").change(function () {
            $(this).find("option").map(function(){
                $("#"+this.value).hide()
            });
            var $option = $(this).find('option:selected');
            $("#"+$option.val()).show()
        });

        $('[data-toggle="tooltip"]').tooltip();
});

function escapeHtml(text) {
    var map = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#039;'
    };

    return text.replace(/[&<>"']/g, function(m) { return map[m]; });
}

function IsJsonString(str) {
    try {
        JSON.parse(str);
    } catch (e) {
        return false;
    }
    return true;
}

String.prototype.replaceAll = function (replaceThis, withThis) {
   var re = new RegExp(RegExp.quote(replaceThis),"g"); 
   return this.replace(re, withThis);
};


RegExp.quote = function(str) {
     return str.replace(/([.?*+^$[\]\\(){}-])/g, "\\$1");
};

function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}


