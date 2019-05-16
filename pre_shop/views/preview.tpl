{{template "tpl/head.html" .}}
 <div class="main">
    <div class="content">
    	<div class="content_top">
    		<div class="back-links">
    		<p><a href="index.html">首页</a >>>>> 商品细节</p>
    	    </div>
    		<div class="clear"></div>
    	</div>
    	<div class="section group">
				<div class="cont-desc span_1_of_2">
				  <div class="product-details">				
					<div class="grid images_3_of_2">
						<div id="container">
						   <div id="products_example">
							   <div id="products">
								<div class="slides_container">
									<a href="#" target="_blank"><img src="/{{.Goods.Image}}" alt=" " /></a>
								</div>
							</div>
						</div>
					</div>
				</div>
				<div class="desc span_3_of_2">
					<h2>{{.Goods.Title}} </h2>
					<p>{{.Goods.Desc}}</p>
					<div class="price">
						<p>Price: <span>{{.Goods.Price}}¥</span></p>
					</div>
					<div class="available">
					<ul>
					    分类：{{.Goods.CategoryName}}
					</ul>
					<ul>
                        卖家：{{.Goods.Username}}
                    </ul>
					</div>
				<div class="share-desc">
					<div class="share">

					</div>
					<div class="button"><span><a href="/goods/cart?id={{.Goods.Id}}&title={{.Goods.Title}}&price={{.Goods.Price}}">Add to Cart</a></span></div>
					<div class="clear"></div>
				</div>
				 <div class="wish-list">
				 	<ul>

				 	</ul>
				 </div>
			</div>
			<div class="clear"></div>
		  </div>
		<div class="product_desc">	
			<div id="horizontalTab">
				<ul class="resp-tabs-list">
					<li>商品细节 & 转卖原因</li>
					<li></li>
					<div class="clear"></div>
				</ul>
				<div class="resp-tabs-container">
					<div class="product-desc">
						商品细节：<p>{{.Goods.Desc}}</p>
						转卖原因：<p>{{.Goods.Cause}}</p>	</div>
						{{if .Goods.Origin_price}}
						<div class="price">
                            <p>原始价格: <span>{{.Goods.Origin_price}}¥</span></p>
                        </div>
                        {{end}}
				 <div class="product-tags">
						 <p></p></div>

				<div class="review">
					<div class="your-review">
				  	 </div>
				</div>
			</div>
		 </div>
	 </div>
	    <script type="text/javascript">
    $(document).ready(function () {
        $('#horizontalTab').easyResponsiveTabs({
            type: 'default', //Types: default, vertical, accordion           
            width: 'auto', //auto or any width like 600px
            fit: true   // 100% fit in a container
        });
    });
   </script>		
   <div class="content_bottom">
    		<div class="heading">
    		<h3>相关商品</h3>
    		</div>
    		<div class="see">
    			<p><a href="/goods/category?category={{.Goods.Category}}">查看所有</a></p>
    		</div>
    		<div class="clear"></div>
    	</div>
   <div class="section group">
				{{range $index, $elem := .Related}}
                   {{if $elem.Id}}
                    <div style="height:280px;" class="grid_1_of_4 images_1_of_4">
                         <a href="preview.html?id={{$elem.Id}}"><img style="height:200px;width:200px;" src="/{{$elem.Image}}" alt="" /></a>
                         <h2>{{$elem.Title}} </h2>
                        <div class="price-details">
                           <div class="price-number">
                                <p><span class="rupees">{{$elem.Price}}¥</span></p>
                            </div>
                                    <div class="add-cart">
                                        <h4><a href="preview.html?id={{$elem.Id}}">Add to Cart</a></h4>
                                     </div>
                                 <div class="clear"></div>
                        </div>
                    </div>
                   {{end}}
                {{end}}

			</div>
        </div>
				<div class="rightsidebar span_3_of_1">
					<h2>商品分类</h2>
					<ul>
				      {{range $index, $elem := .Cat}}
                         {{if $elem}}
                              <li><a href="/goods/category?category={{$index}}"> {{$elem}} </a></li>
                         {{end}}
                     {{end}}
    				</ul>
    				<div class="subscribe">

      				</div>
      				 <div class="community-poll">
      				 	</div>
      				 </div>
 				</div>
 		</div>
 	</div>
    </div>
 </div>
{{template "tpl/foot.html" .}}