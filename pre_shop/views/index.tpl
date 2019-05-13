<!DOCTYPE HTML>

<head>
<title>Home</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<link href="/static/css/style.css" rel="stylesheet" type="text/css" media="all"/>
<link href="/static/css/slider.css" rel="stylesheet" type="text/css" media="all"/>
<script type="text/javascript" src="/static/js/jquery-1.7.2.min.js"></script>
<script type="text/javascript" src="/static/js/move-top.js"></script>
<script type="text/javascript" src="/static/js/easing.js"></script>
<script type="text/javascript" src="/static/js/startstop-slider.js"></script>
</head>
<body>
  <div class="wrap">
	<div class="header">
		<div class="headertop_desc">
			<div class="call">
			  <p><span>需要帮助?</span> 拨打 <span class="number">1-22-3456789</span></span></p>
			</div>
			<div class="account_desc">
				<ul>
					<li><a href="/register">注册</a></li>
                    <li><a href="/login">登录</a></li>
                    <li><a href="delivery.html">物流</a></li>
                    <li><a href="#">结算</a></li>
                    <li><a href="/user/info">个人中心</a></li>
				</ul>
			</div>
			<div class="clear"></div>
		</div>
		<div class="header_top">
			<div class="logo">
				<a href="index.html"><img src="/static/images/logo.png" alt="" /></a>
			</div>
			  <div class="cart">
			  	   <p> <span>购物车:</span><div id="dd" class="wrapper-dropdown-2"> 0 item(s) - ¥0.00
			  	   	<ul class="dropdown">
							<li>没有商品</li>
					</ul></div></p>
			  </div>
			  <script type="text/javascript">
			function DropDown(el) {
				this.dd = el;
				this.initEvents();
			}
			DropDown.prototype = {
				initEvents : function() {
					var obj = this;

					obj.dd.on('click', function(event){
						$(this).toggleClass('active');
						event.stopPropagation();
					});
				}
			}

			$(function() {

				var dd = new DropDown( $('#dd') );

				$(document).click(function() {
					// all dropdowns
					$('.wrapper-dropdown-2').removeClass('active');
				});

			});

		</script>
	 <div class="clear"></div>
  </div>
	<div class="header_bottom">
	     	<div class="menu">
	     		<ul>
			    	<li class="active"><a href="index">首页</a></li>
                    <li><a href="about">关于</a></li>
                    <li><a href="delivery">物流</a></li>
                    <li><a href="news">上新</a></li>
                    <li><a href="contact">联系我们</a></li>
                    <div class="clear"></div>
     			</ul>
	     	</div>
	     	<div class="search_box">
	     		<form>
	     			<input type="text" value="Search" onfocus="this.value = '';" onblur="if (this.value == '') {this.value = 'Search';}"><input type="submit" value="">
	     		</form>
	     	</div>
	     	<div class="clear"></div>
	     </div>
	<div class="header_slide">
			<div class="header_bottom_left">				
				<div class="categories">
				  <ul>
				  	<h3>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;分类</h3>
                       {{range $index, $elem := .Cat}}
                           {{if $elem}}
                                <li><a href="/category/{{$index}}"> {{$elem}} </a></li>
                           {{end}}
                       {{end}}
				  </ul>
				</div>					
	  	     </div>
					 <div class="header_bottom_right">					 
					 	 <div class="slider">					     
							 <div id="slider">
			                    <div id="mover">
			                    	<div id="slide-1" class="slide">
									 <div class="slider-img">
									     <a href="preview.html"><img src="/static/upload/images/ea0402857fd1189fc08251039a1ab4f9.png" alt="learn more" /></a>
									  </div>
						             	<div class="slider-text">
		                                 <h1>精品推荐<br><span>9成新</span></h1>
		                                 <h2>物美价廉</h2>
									   <div class="features_list">
									   	<h4>外观好看的音响</h4>
							            </div>
							             <a href="preview.html" class="button">现在购买</a>
					                   </div>
									  <div class="clear"></div>
				                  </div>	
						             	<div class="slide">
						             		<div class="slider-text">
		                                 <h1>精品推荐<br><span>8成新</span></h1>
		                                 <h2>物美价廉</h2>
									   <div class="features_list">
									   	<h4>大彩电低价出售</h4>
							            </div>
							             <a href="preview.html" class="button">Shop Now</a>
					                   </div>
						             	 <div class="slider-img">
									     <a href="preview.html"><img src="/static/upload/images/6bcd58043eb4d4c856fcb4bce7c0fa16.jpg" alt="learn more" /></a>
									  </div>
									  <div class="clear"></div>
				                  </div>
				                  <div class="slide">
					                  <div class="slider-img">
									     <a href="preview.html"><img src="/static/upload/images/a2d2d2b63a7d88e928b1e6c45f1d7c1c.jpg" alt="learn more" /></a>
									  </div>
									  <div class="slider-text">
		                                 <h1>低价<br><span>出售</span></h1>
		                                 <h2>精美推荐</h2>
									   <div class="features_list">
									   	<h4>超薄27寸电视</h4>
							            </div>
							             <a href="preview.html" class="button">Shop Now</a>
					                   </div>
									  <div class="clear"></div>
				                  </div>
			                 </div>
		                </div>
					 <div class="clear"></div>					       
		         </div>
		      </div>
		   <div class="clear"></div>
		</div>
   </div>
 <div class="main">
    <div class="content">
    	<div class="content_top">
    		<div class="heading">
    		<h3>热卖商品</h3>
    		</div>
    		<div class="see">
    			<p><a href="#">更多商品</a></p>
    		</div>
    		<div class="clear"></div>
    	</div>
	      <div class="section group">
            {{range $index, $elem := .Goods1}}
               {{if $elem.Id}}
				<div class="grid_1_of_4 images_1_of_4">
					 <a href="preview.html"><img style="height:300px;width:300px;" src="/{{$elem.Image}}" alt="" /></a>
					 <h2>{{$elem.Title}} </h2>
					<div class="price-details">
				       <div class="price-number">
							<p><span class="rupees">{{$elem.Price}}¥</span></p>
					    </div>
					       		<div class="add-cart">								
									<h4><a href="preview.html">Add to Cart</a></h4>
							     </div>
							 <div class="clear"></div>
					</div>
				</div>
            {{end}}
            {{end}}

			</div>
			<div class="content_bottom">
    		<div class="heading">
    		<h3>商品上新</h3>
    		</div>
    		<div class="see">
    			<p><a href="#">更多商品</a></p>
    		</div>
    		<div class="clear"></div>
    	</div>

			<div class="section group">
				{{range $index, $elem := .Goods2}}
				{{if $elem.Id}}
                    <div class="grid_1_of_4 images_1_of_4">
                         <a href="preview.html"><img style="height:300px;width:300px;" src="/{{$elem.Image}}" alt="" /></a>
                         <h2>{{$elem.Title}} </h2>
                        <div class="price-details">
                           <div class="price-number">
                                <p><span class="rupees">{{$elem.Price}}¥</span></p>
                            </div>
                                    <div class="add-cart">
                                        <h4><a href="preview.html">Add to Cart</a></h4>
                                     </div>
                                 <div class="clear"></div>
                        </div>
                    </div>
                {{end}}
            {{end}}
			</div>
    </div>
 </div>
</div>
{{template "tpl/foot.html" .}}