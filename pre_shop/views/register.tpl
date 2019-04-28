{{template "tpl/head.html" .}}
<style>
    .main1 {
        margin: 0 auto;
        padding-left: 25px;
        padding-right: 25px;
        padding-top: 15px;
        width: 350px;
        height: 350px;
        background: #FFFFFF;
        /*以下css用于让登录表单垂直居中在界面,可删除*/
        position: absolute;
        top: 50%;
        left: 50%;
        margin-top: -150px;
        margin-left: -175px;
    }

    .title {
        width: 100%;
        height: 40px;
        line-height: 40px;
    }

    .title span {
        font-size: 18px;
        color: #353f42;
    }

    .title-msg {
        width: 100%;
        height: 64px;
        line-height: 64px;
    }

    .title:hover{
        cursor: default	;
    }

    .title-msg:hover{
        cursor: default	;
    }

    .title-msg span {
        font-size: 12px;
        color: #707472;
    }

    .input-content {
        width: 100%;
        height: 120px;
    }

    .input-content input {
        width: 330px;
        height: 40px;
        border: 1px solid #dad9d6;
        background: #ffffff;
        padding-left: 10px;
        padding-right: 10px;
    }

    .enter-btn {
        width: 350px;
        height: 40px;
        color: #fff;
        background: #0bc5de;
        line-height: 40px;
        text-align: center;
        border: 0px;
    }

    .foor{
        width: 100%;
        height: auto;
        color: #9b9c98;
        font-size: 12px;
        margin-top: 20px;
    }

    .enter-btn:hover {
        cursor:pointer;
        background: #1db5c9;
    }

    .foor div:hover {
        cursor:pointer;
        color: #484847;
        font-weight: 600;
    }

    .left{
        float: left;
    }
    .right{
        float: right;
    }

</style>
<div class="main1">
    <div class="title">
        <span>注册账户</span>
    </div>

    <form action="user/register" class="login-form" method="post" >
        <!--输入框-->
        <div class="input-content">
            <!--autoFocus-->
            <div>
                <input type="text" autocomplete="off" placeholder="用户名" name="username" required/>
            </div>

            <div style="margin-top: 16px">
                <input type="text" autocomplete="off" placeholder="手机号" name="phone" required maxlength="32"/>
            </div>
        </div>

        <!--输入框-->
        <div class="input-content">
            <!--autoFocus-->
            <div>
                <input type="password" autocomplete="off" placeholder="登录密码" name="password" required/>
            </div>

            <div style="margin-top: 16px">
                <input type="password" autocomplete="off" placeholder="确认登录密码" name="repassword" required maxlength="32"/>
            </div>
        </div>

        <!--输入框-->
        <div class="input-content">
            <!--autoFocus-->
            <div>
                <input type="email" autocomplete="off" placeholder="邮箱" name="email" required/>
            </div>

            <div style="margin-top: 16px">
                <input type="text" autocomplete="off" placeholder="真实姓名" name="real_name" required maxlength="32"/>
            </div>
        </div>

        <!--登入按钮-->
        <div style="text-align: center">
            <button type="submit" class="enter-btn" >注册</button>
        </div>


    </form>

</div>
