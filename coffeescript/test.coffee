feedUrl = 'http://www.google.com/finance/info?client=ig&q=goog,msft'

@getFeed = ->
    xhr = new XMLHttpRequest
    xhr.open "GET", feedUrl, true
    xhr.onreadystatechange = ->
        if xhr.readyState is 200
            j = JSON.parse xhr.responseText
            x = ""
            j.forEach = (symbol) ->
                x += symbol.t + " "
            $("names").innerHTML = x
    xhr.send null
