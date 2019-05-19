{{template "tpl/head.html" .}}

<style type="text/css">
	/*设置 a 标签的css样式*/
	.page-normal a{
		border: 1px solid #ff6500;
		padding: 5px 7px;
		color: #ff6500;
		margin-left: 20px;/*设置标签 a 之间的间隔*/
		text-decoration: none;/*去除页码数字下面的下划线*/
	}
	/*设置鼠标经过时 a 标签的css样式*/
	.page-normal a:hover{
		background-color: #ffbe94;
	}
	/*设置整个div的css样式,该样式主要用于设置网页内的省略号 …… 的样式，并且同时设置内容居中显示*/
	.page-normal{
		color: #ff6500;
		text-align: center;
	}
	/*设置当前页面的css样式*/
	.page-normal .page-current{
		color: #ffffff;
		background-color: #ff6500;
	}
	/*进行代码优化，将不同css样式中共有的属性放在一起，有助于提高运行效率*/
	.page-normal a, .page-normal a:hover, .page-normal .page-prev, .page-normal .page-current{
		border: 1px solid #ff6500;
		padding: 5px 7px;
	}
</style>

 <div class="main">
    <div class="content">
    	<div class="content_top">
    		<div class="heading">
    		<p><a href="/index">首页</a> >>>> 商品列表</p>
    		</div>
    		<div class="see">

    		</div>
    		<div class="clear"></div>
    	</div>
	      <div class="section group">
				{{range $index, $elem := .Goods1}}
                   {{if $elem.Id}}
                    <div class="grid_1_of_4 images_1_of_4">
                         <a href="/preview.html?id={{$elem.Id}}"><img style="height:300px;width:300px;" src="/{{$elem.Image}}" alt="" /></a>
                         <h2>{{$elem.Title}} </h2>
                        <div class="price-details">
                           <div class="price-number">
                                <p><span class="rupees">{{$elem.Price}}¥</span></p>
                            </div>
                                    <div class="add-cart">
                                        <h4><a href="/preview.html?id={{$elem.Id}}">Add to Cart</a></h4>
                                     </div>
                                 <div class="clear"></div>
                        </div>
                    </div>
                {{end}}
                {{end}}
			</div>

        <div class="section group">
            {{range $index, $elem := .Goods2}}
               {{if $elem.Id}}
                <div class="grid_1_of_4 images_1_of_4">
                     <a href="/preview.html?id={{$elem.Id}}"><img style="height:300px;width:300px;" src="/{{$elem.Image}}" alt="" /></a>
                     <h2>{{$elem.Title}} </h2>
                    <div class="price-details">
                       <div class="price-number">
                            <p><span class="rupees">{{$elem.Price}}¥</span></p>
                        </div>
                                <div class="add-cart">
                                    <h4><a href="/preview.html?id={{$elem.Id}}">Add to Cart</a></h4>
                                 </div>
                             <div class="clear"></div>
                    </div>
                </div>
            {{end}}
            {{end}}
        </div>

        <div class="section group">
            {{range $index, $elem := .Goods3}}
               {{if $elem.Id}}
                <div class="grid_1_of_4 images_1_of_4">
                     <a href="/preview.html?id={{$elem.Id}}"><img style="height:300px;width:300px;" src="/{{$elem.Image}}" alt="" /></a>
                     <h2>{{$elem.Title}} </h2>
                    <div class="price-details">
                       <div class="price-number">
                            <p><span class="rupees">{{$elem.Price}}¥</span></p>
                        </div>
                                <div class="add-cart">
                                    <h4><a href="/preview.html?id={{$elem.Id}}">Add to Cart</a></h4>
                                 </div>
                             <div class="clear"></div>
                    </div>
                </div>
            {{end}}
            {{end}}
        </div>
    </div>
 </div>

<br>
<div class="page-normal">
	<a href="/goods/list?page={{.Prev}}">上一页</a><a href="/goods/list?page={{.Next}}">下一页</a>
	&nbsp;&nbsp;&nbsp;&nbsp;<span>共{{.Count}}条数据,当前第{{.Page}}页</span>

</div>



{{template "tpl/foot.html" .}}