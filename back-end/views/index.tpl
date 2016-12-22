<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="referrer" content="origin">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/css/news.css">
    <title>Hacker News</title></head>
<body>
<center>
    <table id="hnmain" border="0" cellpadding="0" cellspacing="0" width="85%" bgcolor="#f6f6ef">
        <tbody>
        <tr>
            <td bgcolor="#ff6600">
                <table border="0" cellpadding="0" cellspacing="0" width="100%" style="padding:2px">
                    <tbody>
                    <tr>
                        <td style="width:18px;padding-right:4px"><a href="http://www.ycombinator.com/"><img
                                src="/static/y18.gif" width="18" height="18"
                                style="border:1px white solid;"></a></td>
                        <td style="line-height:12pt; height:10px;"><span class="pagetop"><b class="hnname"><a href="https://news.ycombinator.com/news">Hacker News</a></b>
                        </span></td>
                        <td style="text-align:right;padding-right:4px;"><span class="pagetop">
                              <a href="/login">login</a>
                          </span></td>
                    </tr>
                    </tbody>
                </table>
            </td>
        </tr>
        <tr style="height:10px"></tr>
        <tr>
            <td>
                <table border="0" cellpadding="0" cellspacing="0" class="itemlist">
                    <tbody>
                    {{range $key,$response := .liste}}
                    <tr class="athing">
                        <td align="right" valign="top" class="title"></td>
                        <td valign="top" class="votelinks">
                        </td>
                        <td class="title"><a href="#" class="storylink">{{$response.News.Title}}</a></td>
                    </tr>
                    <tr>
                        <td colspan="2"></td>
                        <td class="subtext">{{$response.News.Content}}
                            <span class="age"></span> <span id="unv_13239577"></span> |
                            <a href="#">{{$response.User.Firstname}} </a></td>
                    </tr>
                    <tr class="spacer" style="height:5px"></tr>
                    {{end}}
                    </tbody>
                </table>
            </td>
        </tr>
        <tr>
            <td><img src="/static/s.gif" height="10" width="0">
                <table width="100%" cellspacing="0" cellpadding="1">
                    <tbody>
                    <tr>
                        <td bgcolor="#ff6600"></td>
                    </tr>
                    </tbody>
                </table>
                <br>
            </td>
        </tr>
        </tbody>
    </table>
</center>
</body>
</html>