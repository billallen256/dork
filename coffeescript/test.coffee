feedUrl = 'http://api.openweathermap.org/data/2.5/weather?lat=39&lon=-104'

$ = (id) -> document.getElementById(id)

@getFeed = ->
    xhr = new XMLHttpRequest
    xhr.open "GET", feedUrl, true
    xhr.onreadystatechange = ->
        if xhr.readyState is 4
            if xhr.status is 200
                j = JSON.parse xhr.responseText
                x = "<table>"
                for i of j.weather
                    for k,v of i
                        x += "<tr><td>"+k+"</td><td>"+v+"</td></tr>"
                x += "</table>"
                $("names").innerHTML = x
    xhr.send null
