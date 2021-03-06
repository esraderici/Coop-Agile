<html>
<head>
    <style>
        .main {

            width:40%;
            margin:auto;
        }

        .form-group {
            width:90%;
            padding:5%;
            padding-bottom:1%;
        }
        .form-group input{
            height:20px;
            width:100%;
            font-size:17px;
            padding:15px;
            border-radius:10px;
        }

        .form-group button{
            border-radius:10px;
            width:48%;
            height:40px;
            background-color:#28c31d;
            font-size:15px;
            color:white;
            border:0px;
        }

        .form-group button:hover{
            background-color:gray
        }

        .registration tr td{

            font-family:'Tahoma';

        }
    </style>
</head>
<body>
<div class="main">

    <form method="post" action="/user/register">
        <table class="registration">
            <tr>
                <td colspan="2">
                    <div class="form-group">
                        <h3 style="font-family:'Tahoma';text-align:center;color: #717171;">SIGN UP</h3>
                    </div>
                </td>
            </tr>
            <tr>
                <td></td>
                <td>
                    <div class="form-group">
                     {{.error}}
                    </div>
                </td>
            </tr>
            <tr>
                <td>First Name</td>
                <td>
                    <div class="form-group">
                        <input name="firstname" type="text" />
                    </div>
                </td>
            </tr>
            <tr>
                <td>Last Name</td>
                <td>
                    <div class="form-group">
                        <input name="lastname" type="text" />
                    </div>
                </td>
            </tr>
            <tr>
                <td>Email</td>
                <td>
                    <div class="form-group">
                        <input name="email" type="email" />
                    </div>
                </td>
            </tr>
            <tr>
                <td>Username</td>
                <td>
                    <div class="form-group">
                        <input name="username" type="text" />
                    </div>
                </td>
            </tr>
            <tr>
                <td>Password</td>
                <td>
                    <div class="form-group">
                        <input name="password" type="password" />
                    </div>
                </td>
            </tr>
            <tr>
                <td colspan="2">
                    <div class="form-group" style="text-align:right;" >
                        <button type="submit" style="">Create</button>
                    </div>
                </td>
            </tr>
            <tr>
                <td colspan="2">
                    <div class="form-group" style="text-align:right;" >
                        <a href="/login" style="color:#00BCD4;background-color: white; font-size:15px;">You have already an account?</a>
                    </div>
                </td>
            </tr>

        </table>
    </form>

</div>
</body>
</html>