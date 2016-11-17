<html>
<head>
	<style>
		.main {

			width:30%;
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
				width:60%;
				height:40px;
				background-color:#28c31d;
				font-size:15px;
				color:white;
				border:0px;
			}

			.form-group button:hover{
				background-color:gray
			}
	</style>
</head>

<body>
	<div class="main">
		<form method="post" >
		{{if  .error }}
		<div class="form-group">
        				Error : {{.error}}
        			</div>

        {{end}}
			<div class="form-group">
				<h3 style="font-family:'Tahoma';text-align:center;color: #717171;">SIGN IN</h3>
			</div>
			<div class="form-group">
				<input type="text" placeholder="Username" name="username"/>
			</div>
			<div class="form-group">
				<input type="password" placeholder="Password" name="password"/>
			</div>
			<div class="form-group" style="text-align:center;" >
				<button type="submit" style="">LOGIN</button>
			</div>
			<div class="form-group" style="text-align:center;" >
				<a href="/user/register" style="color:#00BCD4;background-color: white;">Don't you have an account?</a>
			</div>
		</form>
	</div>
</body>
</html>