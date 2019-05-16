{{template "tpl/head.html" .}}
 <div class="main">
    <div class="content">
    	<div class="section group">
				<div class="col span_2_of_3">
				  <div class="contact-form">
				  	<h2>Contact Us</h2>
					    <form method="get" action="/about/contact">
					    	<div>
						    	<span><label>Name</label></span>
						    	<span><input name="userName" type="text" class="textbox" ></span>
						    </div>
						    <div>
						    	<span><label>E-mail</label></span>
						    	<span><input name="userEmail" type="text" class="textbox"></span>
						    </div>
						    <div>
						     	<span><label>Company Name</label></span>
						    	<span><input name="userPhone" type="text" class="textbox"></span>
						    </div>
						    <div>
						    	<span><label>Subject</label></span>
						    	<span><textarea name="userMsg"> </textarea></span>
						    </div>
						   <div>
						   		<span><input type="submit" value="Submit"  class="myButton"></span>
						  </div>
					    </form>
				  </div>
  				</div>
				<div class="col span_1_of_3">

      			<div class="company_address">
				     	<h3>开发者信息 :</h3>
						    	<p>联系方式：13148241118</p>
						    	<p>地址：河南省南阳市长江路八十号南阳理工学院</p>
				   		<p>Phone:13148241119</p>
				 	 	<p>Email: <span>1135030879@qq.com</span></p>
				   </div>
				 </div>
			  </div>		
         </div> 
    </div>
 </div>
{{template "tpl/foot.html" .}}

