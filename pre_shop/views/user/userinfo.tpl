<html>
<head>
    <meta charset="utf-8">
    <title>{{.Title}}</title>
    <script type="text/javascript" src="/static/user/js/safe/aywmq_qt.js"></script>
    <script type="text/javascript" src="/static/user/js/safe/da_opt.js"></script>
    <meta name="keywords" content="">
    <meta name="description" content="">
    <link rel="stylesheet" href="/static/user/css/safe/css.css" />
    <link rel="stylesheet" href="/static/user/css/safe/common.min.css" />
    <link rel="stylesheet" href="/static/user/css/safe/ms-style.min.css" />
    <link rel="stylesheet" href="/static/user/css/safe/personal_member.min.css" />
    <link rel="stylesheet" href="/static/user/css/safe/Snaddress.min.css" />
    <link rel="stylesheet" href="/static/user/css/sui.css" />
    <script type="text/javascript" src="/static/user/js/jquery-1.9.1.min.js" ></script>
    <script type="text/javascript" src="/static/user/js/sui.js" ></script>
    <style>
            .fenye{
                border: 1px #ccc solid;
                 padding: 4px;
                width: 20px;

        }
        .sui-nav.nav-tabs {
            border-bottom: 1px solid #F88600;
         padding-left: 0px;
    }
    </style>
</head>

<body class="ms-body">
    <div id="" class="ng-top-banner"></div>
    <div class="ng-toolbar">
        <div class="ng-toolbar-con wrapper">
            <div class="ng-toolbar-left">
                <a href="/" class="ng-bar-node ng-bar-node-backhome" id="ng-bar-node-backhome" style="display: block;">
                    <span><img src="/static/user/img/Home.png" style="margin-right: 10px;"/>返回首页</span>
                </a>
            <div class="ng-toolbar-right">
                </div>
            </div>
            <div id="ng-minicart-slide-box"></div>
        </div>
    </div>
    <header class="ms-header ms-header-inner ms-head-position">
        <article class="ms-header-menu">
            <style type="text/css">
                .nav-manage .list-nav-manage {
                    position: absolute;
                    padding: 15px 4px 10px 15px;
                    left: 0;
                    top: -15px;
                    width: 90px;
                    background: #FFF;
                    box-shadow: 1px 1px 2px #e3e3e3, -1px 1px 2px #e3e3e3;
                    z-index: 10;
                }

                .ms-nav li {
                    float: left;
                    position: relative;
                    padding: 0 20px;
                    height: 44px;
                    font: 14px/26px "Microsoft YaHei";
                    color: #FFF;
                    cursor: pointer;
                    z-index: 10;
                }
            </style>
            <div class="header-menu">
                <div class="ms-logo">
                    <a class="ms-head-logo" name="Myyigou_index_none_daohangLogo"><span style="font-size: 30px;color: #fff;font-weight: bold;    line-height: 28px;;"></span></a>

                </div>
            </div>

        </article>

        <article class="ms-useinfo">
            <div class="header-useinfo" id="">
                <div class="ms-avatar">
                    <div class="useinfo-avatar">
                        <img src="/static/user/img/头像.png">
                        <div class="edit-avatar"></div>
                        <a class="text-edit-avatar">修改</a>
                    </div>
                    <a>{{.Username}}</a>
                </div>
                <div class="ms-name-info">
                    <div class="info-safety">
                        <a class="manage-addr"><i style="background-image: url(/static/user/img/地址管理.png);"></i>地址管理</a>
                    </div>
                </div>
            </div>

        </article>
    </header>
    <div id="ms-center" class="personal-member">
        <div class="cont">
            <div class="cont-side">
                <div class="side-neck" style="margin-top: 20px;">
                    <i></i>
                </div>
                <div class="ms-side" style="margin-top: 20px;">
                    <article class="side-menu side-menu-off">
                        <dl class="side-menu-tree" style="padding-left: 30px;">
                            <dt><img src="/static/user/img/左侧/file.png"  style="margin-right: 10px;margin-left: -20px;"/>作为卖家</dt>
                            <dd>
                                <a href="#profile" data-toggle="tab">我的发售</a>
                            </dd>
                            <dd>
                                <a href="#tixing" data-toggle="tab">发售商品</a>
                            </dd>
                            <dd>
                                <a>商品评价</a>
                            </dd>
                            <dt><img src="/static/user/img/左侧/我的买啦.png"  style="margin-right: 10px;margin-left: -20px;"/>作为买家</dt>
                            <dd>
                                <a>我的购买</a>
                            </dd>
                            <dd>
                                <a>我的订单</a>
                            </dd>
                            <dt><img src="/static/user/img/左侧/v-card-3.png"  style="margin-right: 10px;margin-left: -20px;"/>其他</dt>
                            <dd>
                                <a >我的个人信息</a>
                            </dd>
                            <dd>
                                <a>我的地址</a>
                            </dd>
                        </dl>

                        <a ison="on" class="switch-side-menu icon-up-side"><i></i></a>
                    </article>
                </div>
            </div>
            <div class="cont-main">
                <div class="main-wrap mt15" style="border: 0px;">
                    <div class="server-wrapper">
                        <div class="server-tab" style="margin-top: 20px;">

                            <ul class="sui-nav nav-tabs">
                              <li class="active"><a href="#index" data-toggle="tab">我的提醒</a></li>
                              <li><a href="#profile" data-toggle="tab">我的发售</a></li>
                              <li><a href="#about" data-toggle="tab">我的购买</a></li>
                              <li><a href="#tixing" data-toggle="tab">发售商品</a></li>
                              <li><a href="#set" data-toggle="tab">我的地址</a></li>
                            </ul>
                            <div class="tab-content">
                             <div id="index" class="tab-pane active">
                                    <table class="uiTable" id="return_table"  style="margin-top: 20px;border:  1px #ccc solid;    width: 90%;">
                                        <tr  style="background-color: #f8f8f8;" >
                                            <th  style="line-height:50px;margin-left:0px;" nowrap><font style="font-size: 14px; ">通知公告</font></th>
                                            <th  style="line-height:50px;margin-left:0px;" nowrap><font style="font-size: 14px; ">通知公告</font></th>
                                            <th  style="line-height:50px;" nowrap><font style="font-size: 14px;margin-left: 41px; ">在售商品</font></th>
                                            <th  style="line-height:50px; "nowrap ><font style="font-size: 14px;margin-left:28px; ">在线订单</font></th>
                                        </tr>
                                        <tr >
                                               <td style="font-size: 12px;">
                                                   <span style="margin-left: 41px;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                                    500
                                                   </span>
                                               </td>
                                               <td style="font-size: 12px">
                                                   <span style="margin-left:40px;">0</span>
                                               </td>
                                               <td style="font-size: 12px">
                                                       <span style="margin-left: 40px;">0</span>
                                               </td>
                                               <td style="font-size: 12px">
                                                   <span style="margin-left: 28px;">5300.00</span>
                                               </td>
                                               <td style="font-size: 12px"></td>
                                        </tr>
                                    </table>
                              </div>
                              <div id="profile" class="tab-pane">
                                <table class="uiTable" id="return_table"  style="margin-top: 20px;border:  1px #ccc solid;    width: 90%;">
                                    <tr  style="background-color: #f8f8f8;" >
                                        <th  style="line-height:50px;" nowrap colspan="2"><font style="font-size: 14px;margin-left:90px; ">商品名称</font></th>
                                        <th  style="line-height:50px;" nowrap><font style="font-size: 14px;margin-left: 41px; ">商品价格</font></th>
                                        <th  style="line-height:50px; "nowrap ><font style="font-size: 14px;margin-left:28px; ">商品简介</font></th>
                                        <th  style="line-height:50px; "  nowrap><font style="font-size: 14px;margin-left:35px; ">商品分类</font></th>
                                        <th  style="line-height:50px; " nowrap ><font style="font-size: 14px;margin-left:26px; ">商品转卖原因</font></th>
                                        <th  style="line-height:50px; " nowrap><font style="font-size: 14px;margin-left:34px; ">发售时间</font></th>
                                        <th  style="line-height:50px; " nowrap ><font style="font-size: 14px;margin-left:34px;margin-right: 20px; ">状态</font></th>
                                    </tr>
                                    {{range $index, $elem := .Goods}}
                                      {{if $elem.Id}}
                                        <tr >
                                            <td style="padding: 20px;">
                                            <img src="/{{$elem.Image}}" style="max-width: 80px;max-height: 80px;margin-left: ;" />
                                            </td>
                                            <td >
                                                <p style="color: #284ec7;font-size: 12px;width: 100px;margin-top:0px;">{{$elem.Title}}</p>
                                            </td>
                                            <td style="font-size: 12px;">
                                                <span style="margin-left: 41px;">{{$elem.Price}}
                                                </span>
                                            </td>
                                            <td style="font-size: 12px">
                                                <span style="margin-left:40px;">{{$elem.Desc}}</span>
                                            </td>
                                            <td style="font-size: 12px">
                                                    <span style="margin-left: 40px;">{{$elem.CategoryName}}</span>
                                            </td>
                                            <td style="font-size: 12px">
                                                <span style="margin-left: 28px;">{{$elem.Cause}}</span>
                                            </td>
                                            <td style="font-size: 12px">{{$elem.Add_time}}</td>
                                            <td style="font-size: 12px">

                                                <span style="margin-left: 34px;">{{$elem.Price}}</span>
                                                <br />
                                                <br />
                                             </td>
                                        </tr>
                                     {{end}}
                                  {{end}}
                                </table>
                              </div>
                              <div id="about" class="tab-pane">
                                <table class="uiTable" id="return_table"  style="margin-top: 20px;border:  1px #ccc solid;    width: 90%;">
                                    <tr  style="background-color: #f8f8f8;" >
                                        <th  style="line-height:50px;" nowrap colspan="2"><font style="font-size: 14px;margin-left:90px; ">宝贝名称</font></th>
                                        <th  style="line-height:50px;margin-left:135px;" nowrap><font style="font-size: 14px; ">开拍时间</font></th>
                                        <th  style="line-height:50px;" nowrap><font style="font-size: 14px;margin-left: 41px; ">保证金</font></th>
                                        <th  style="line-height:50px; "nowrap ><font style="font-size: 14px;margin-left:28px; ">竞拍数量</font></th>
                                        <th  style="line-height:50px; "  nowrap><font style="font-size: 14px;margin-left:35px; ">获得数量</font></th>
                                        <th  style="line-height:50px; " nowrap ><font style="font-size: 14px;margin-left:26px; ">当前价格(元)</font></th>
                                        <th  style="line-height:50px; " nowrap><font style="font-size: 14px;margin-left:34px; ">我的最高出价(元)</font></th>
                                        <th  style="line-height:50px; " nowrap><font style="font-size: 14px;margin-left:34px; ">剩余时间</font></th>
                                        <th  style="line-height:50px; " nowrap ><font style="font-size: 14px;margin-left:34px;margin-right: 20px; ">状态</font></th>
                                    </tr>
                                    <tr >
                                        <td style="padding: 20px;">
                                        <img src="/static/user/img/我的拍卖/组-43.png" style="max-width: 100px;margin-left: ;" />
                                        </td>
                                        <td >
                                            <p style="color: #284ec7;font-size: 12px;width: 100px;margin-top:0px;">冰雪翡翠花手镯59mm</p>
                                        </td>
                                        <td style="font-size: 12px">
                                            2016-04-15
                                            <br />
                                            10:00:00
                                        </td>
                                        <td style="font-size: 12px;">
                                            <span style="margin-left: 41px;">500
                                            <br /><span style="color: #F2873B;margin-left: 35px;"> 已支付</span>
                                            </span>
                                        </td>
                                        <td style="font-size: 12px">
                                            <span style="margin-left:40px;">0</span>
                                        </td>
                                        <td style="font-size: 12px">
                                                <span style="margin-left: 40px;">0</span>
                                        </td>
                                        <td style="font-size: 12px">
                                            <span style="margin-left: 28px;">5300.00</span>
                                        </td>
                                        <td style="font-size: 12px"></td>
                                        <td style="font-size: 12px">
                                                <span style="margin-left: 34px;">3小时</span></td>
                                        <td style="font-size: 12px">

                                            <span style="margin-left: 34px;">未报名</span>
                                            <br />
                                            <br />
                                        <button style="font-size: 12px;margin-left: 34px;padding: 5px;border: 1px #ccc solid;">报名</button>
                                        </td>
                                    </tr>
                                </table>
                              </div>
                               <div id="tixing" class="tab-pane">
                                <p></p>
                                <form action="/user/goods" method="post" enctype="multipart/form-data">
                                <div class="user-profile-wrap" style="width: 100%;height: 500px;">
                                    <p style="margin-left: 70px;font-size: 14px;"><span style="color:#F88600;font-size: 14px;">发售二手物品</span><span style="margin-left: 10px;font-size: 14px;font-weight: bold;"></span> </p>
                                <div style="margin-left: 70px;margin-top: 30px;height: 30px;">
                                    <span style="color: #F2873B;">*&nbsp;</span><span class="titles">商品标题:</span>
                                    <input type="text" name="title" placeholder="标题 品类品牌型号都是买家喜欢搜索的" style="padding: 5px;width: 300px;margin-left: 14px;" />

                                </div>
                                <div style="margin-left: 70px;margin-top: 50px;">
                                    <span style="color: #F2873B;">*&nbsp;</span><span class="titles">转卖原因:</span><br>
                                </div>
                                <div style="margin-left: 150px;margin-top:-40px;">
                                    <textarea name="cause" style="width:500px;height: 90px;padding: 5px;" placeholder="描述物品的转卖原因，入手渠道和使用感受:"></textarea>
                                </div>
                                <div style="margin-left: 80px;margin-top: 20px;">
                                    <span class="titles">商品照片:</span>
                                    <input type="file" name="images1" style="padding: 5px;width: 300px;margin-left: 14px;" />
                                </div>
                                <div style="margin-left: 80px;margin-top: 20px;">
                                    <span style="color: #F2873B;">*&nbsp;</span><span class="titles">一口价:</span>
                                    <input type="text" name="price" placeholder="您想以什么价位出手" style="padding: 5px;width: 300px;margin-left: 14px;" />
                                </div>
                                <div style="margin-left: 80px;margin-top: 20px;">
                                    <span style="color: #F2873B;">*&nbsp;</span><span class="titles">商品原始价:</span>
                                    <input type="text" name="origin_price" placeholder="新品价格 作为参考" style="padding: 5px;width: 300px;margin-left: 14px;" />
                                </div>
                                <div style="margin-left: 80px;margin-top: 20px;">
                                    <span class="titles">商品描述:</span>
                                    <textarea name="desc" style="width:500px;height: 90px;padding: 5px;" placeholder="商品的描述：包括商品颜色，尺寸等等:"></textarea>
                                    </div>
                                <div style="margin-left: 55px;margin-top: 10px;">
                                    <span style="color: #F2873B;">*</span>
                                    <span class="titles">商品分类:</span>
                                    <select name="category" style="padding: 5px;margin-left: 14px;">
                                        {{range $index, $elem := .Cat}}
                                           {{if $elem}}
                                                <li><a href="/category/{{$index}}"> {{$elem}} </a></li>
                                                <option value="{{$index}}" >{{$elem}}</option>
                                           {{end}}
                                       {{end}}
                                    </select>
                                    </div>
                                <div style="margin-left: 80px;margin-top: 10px;">
                                    <span style="color: #F2873B;">*&nbsp;</span><span class="titles">交易方式:</span>
                                    <select name="deal" style="padding: 5px;margin-left: 14px;">
                                        <option value="1">自提</option>
                                        <option value="2">同城面交</option>
                                        <option value="3">邮寄</option>
                                        <option value="0">三种方式都支持</option>
                                    </select>
                                </div>
                                <button type="submit" style="margin-left:150px;margin-top:10px;background-color:#F37B1D ;color: #fff;width: 100px;height: 30px;border: 0px;border-radius: 5px;">确定</button>
                                </form>
                                </div>
                              </div>

                               <div id="set" class="tab-pane">
                                <table class="uiTable" id="return_table"  style="margin-top: 20px;border:  1px #ccc solid;    width: 90%;">
                                    <tr  style="background-color: #f8f8f8;" >
                                        <th  style="line-height:50px;" nowrap colspan="2"><font style="font-size: 14px;margin-left:90px; ">宝贝名称</font></th>
                                        <th  style="line-height:50px;margin-left:135px;" nowrap><font style="font-size: 14px; ">开拍时间</font></th>
                                        <th  style="line-height:50px;" nowrap><font style="font-size: 14px;margin-left: 41px; ">保证金</font></th>
                                        <th  style="line-height:50px; "nowrap ><font style="font-size: 14px;margin-left:28px; ">竞拍数量</font></th>
                                        <th  style="line-height:50px; "  nowrap><font style="font-size: 14px;margin-left:35px; ">获得数量</font></th>
                                        <th  style="line-height:50px; " nowrap ><font style="font-size: 14px;margin-left:26px; ">当前价格(元)</font></th>
                                        <th  style="line-height:50px; " nowrap><font style="font-size: 14px;margin-left:34px; ">我的最高出价(元)</font></th>
                                        <th  style="line-height:50px; " nowrap><font style="font-size: 14px;margin-left:34px; ">剩余时间</font></th>
                                        <th  style="line-height:50px; " nowrap ><font style="font-size: 14px;margin-left:34px;margin-right: 20px; ">状态</font></th>
                                    </tr>
                                    <tr >
                                        <td style="padding: 20px;">
                                        <img src="/static/user/img/我的拍卖/组-43.png" style="max-width: 100px;margin-left: ;" />
                                        </td>
                                        <td >
                                            <p style="color: #284ec7;font-size: 12px;width: 100px;margin-top:0px;">冰雪翡翠花手镯59mm</p>
                                        </td>
                                        <td style="font-size: 12px">
                                            2016-04-15
                                            <br />
                                            10:00:00
                                        </td>
                                        <td style="font-size: 12px;">
                                            <span style="margin-left: 41px;">500
                                            <br /><span style="color: #F2873B;margin-left: 35px;"> 已支付</span>
                                            </span>
                                        </td>
                                        <td style="font-size: 12px">
                                            <span style="margin-left:40px;">0</span>
                                        </td>
                                        <td style="font-size: 12px">
                                                <span style="margin-left: 40px;">0</span>
                                        </td>
                                        <td style="font-size: 12px">
                                            <span style="margin-left: 28px;">5300.00</span>
                                        </td>
                                        <td style="font-size: 12px"></td>
                                        <td style="font-size: 12px">
                                                <span style="margin-left: 34px;">3小时</span></td>
                                        <td style="font-size: 12px">

                                            <span style="margin-left: 34px;">未报名</span>
                                            <br />
                                            <br />
                                        <button style="font-size: 12px;margin-left: 34px;padding: 5px;border: 1px #ccc solid;">报名</button>
                                        </td>
                                    </tr>
                                </table>
                              </div>
                        </div>
                    </div>

    <style type="text/css ">
        .ng-footer {
            height: 130px;
            margin-top: 0;
        }


        .ng-s-footer {
            height: 130px;
            background: none;
            text-align: center;
        }

        .ng-s-footer p.ng-url-list {
            height: 25px;
            line-height: 25px;
        }

        .ng-s-footer p.ng-url-list a {
            color: #666666;
        }

        .ng-s-footer p.ng-url-list a:hover {
            color: #f60;
        }

        .ng-s-footer .ng-authentication {
            float: none;
            margin: 0 auto;
            height: 25px;
            width: 990px;
            margin-top: 5px;
        }

        .ng-s-footer p.ng-copyright {
            float: none;
            width: 100%;
        }

        .root1200 .ng-s-footer p.ng-copyright {
            width: 100%;
        }
    </style>
    <script type="text/javascript " src="/static/user/js/safe/ms_common.min.js "></script>
</body>

</html>