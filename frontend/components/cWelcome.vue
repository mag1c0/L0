<script lang="ts" setup>
if (process.browser) {
  useHead({
    script: [{children: '!function(){"use strict";var t,i={screen:{elem:null,callback:null,ctx:null,width:0,height:0,left:0,top:0,init:function(t,i,s){return this.elem=document.getElementById(t),this.callback=i||null,"CANVAS"==this.elem.tagName&&(this.ctx=this.elem.getContext("2d")),window.addEventListener("resize",function(){this.resize()}.bind(this),!1),this.elem.onselectstart=function(){return!1},this.elem.ondrag=function(){return!1},s&&this.resize(),this},resize:function(){var t=this.elem;for(this.width=t.offsetWidth,this.height=t.offsetHeight,this.left=0,this.top=0;null!=t;t=t.offsetParent)this.left+=t.offsetLeft,this.top+=t.offsetTop;this.ctx&&(this.elem.width=this.width,this.elem.height=this.height),this.callback&&this.callback()}}},s=function(t,i){this.x=t,this.y=i,this.magnitude=t*t+i*i,this.computed=0,this.force=0};s.prototype.add=function(t){return new s(this.x+t.x,this.y+t.y)};var h=function(t){this.vel=new s((Math.random()>.5?1:-1)*(.2+.25*Math.random()),(Math.random()>.5?1:-1)*(.2+Math.random())),this.pos=new s(.2*t.width+Math.random()*t.width*.6,.2*t.height+Math.random()*t.height*.6),this.size=t.wh/15+(1.4*Math.random()+.1)*(t.wh/15),this.width=t.width,this.height=t.height};h.prototype.move=function(){this.pos.x>=this.width-this.size?(this.vel.x>0&&(this.vel.x=-this.vel.x),this.pos.x=this.width-this.size):this.pos.x<=this.size&&(this.vel.x<0&&(this.vel.x=-this.vel.x),this.pos.x=this.size),this.pos.y>=this.height-this.size?(this.vel.y>0&&(this.vel.y=-this.vel.y),this.pos.y=this.height-this.size):this.pos.y<=this.size&&(this.vel.y<0&&(this.vel.y=-this.vel.y),this.pos.y=this.size),this.pos=this.pos.add(this.vel)};var e=function(t,i,e,r,o){this.step=5,this.width=t,this.height=i,this.wh=Math.min(t,i),this.sx=Math.floor(this.width/this.step),this.sy=Math.floor(this.height/this.step),this.paint=!1,this.metaFill=a(t,i,t,r,o),this.plx=[0,0,1,0,1,1,1,1,1,1,0,1,0,0,0,0],this.ply=[0,0,0,0,0,0,1,0,0,1,1,1,0,1,0,1],this.mscases=[0,3,0,3,1,3,0,3,2,2,0,2,1,1,0],this.ix=[1,0,-1,0,0,1,0,-1,-1,0,1,0,0,1,1,0,0,0,1,1],this.grid=[],this.balls=[],this.iter=0,this.sign=1;for(var n=0;n<(this.sx+2)*(this.sy+2);n++)this.grid[n]=new s(n%(this.sx+2)*this.step,Math.floor(n/(this.sx+2))*this.step);for(var l=0;l<e;l++)this.balls[l]=new h(this)};e.prototype.computeForce=function(t,i,s){var h,e=s||t+i*(this.sx+2);if(0===t||0===i||t===this.sx||i===this.sy)h=.6*this.sign;else{h=0;for(var a,r=this.grid[e],o=0;a=this.balls[o++];)h+=a.size*a.size/(-2*r.x*a.pos.x-2*r.y*a.pos.y+a.pos.magnitude+r.magnitude);h*=this.sign}return this.grid[e].force=h,h},e.prototype.marchingSquares=function(t){var i=t[0],s=t[1],h=t[2],e=i+s*(this.sx+2);if(this.grid[e].computed===this.iter)return!1;for(var a,r=0,o=0;o<4;o++){var l=i+this.ix[o+12]+(s+this.ix[o+16])*(this.sx+2),c=this.grid[l].force;(c>0&&this.sign<0||c<0&&this.sign>0||!c)&&(c=this.computeForce(i+this.ix[o+12],s+this.ix[o+16],l)),Math.abs(c)>1&&(r+=Math.pow(2,o))}if(15===r)return[i,s-1,!1];5===r?a=2===h?3:1:10===r?a=3===h?0:2:(a=this.mscases[r],this.grid[e].computed=this.iter);var d=this.step/(Math.abs(Math.abs(this.grid[i+this.plx[4*a+2]+(s+this.ply[4*a+2])*(this.sx+2)].force)-1)/Math.abs(Math.abs(this.grid[i+this.plx[4*a+3]+(s+this.ply[4*a+3])*(this.sx+2)].force)-1)+1);return n.lineTo(this.grid[i+this.plx[4*a]+(s+this.ply[4*a])*(this.sx+2)].x+this.ix[a]*d,this.grid[i+this.plx[4*a+1]+(s+this.ply[4*a+1])*(this.sx+2)].y+this.ix[a+4]*d),this.paint=!0,[i+this.ix[a+4],s+this.ix[a+8],a]},e.prototype.renderMetaballs=function(){for(var t,i=0;t=this.balls[i++];)t.move();for(this.iter++,this.sign=-this.sign,this.paint=!1,n.fillStyle=this.metaFill,n.beginPath(),i=0;t=this.balls[i++];){var s=[Math.round(t.pos.x/this.step),Math.round(t.pos.y/this.step),!1];do{s=this.marchingSquares(s)}while(s);this.paint&&(n.fill(),n.closePath(),n.beginPath(),this.paint=!1)}};var a=function(t,i,s,h,e){var a=n.createRadialGradient(t/1,i/1,0,t/1,i/1,s);return a.addColorStop(0,h),a.addColorStop(1,e),a},r=function(){requestAnimationFrame(r),n.clearRect(0,0,o.width,o.height),t.renderMetaballs()},o=i.screen.init("bubble",null,!0),n=o.ctx;o.resize(),t=new e(o.width,o.height,6,"#B65CFD","#CD91FF"),r()}();let tabs=document.querySelectorAll(".it-block__tab");for(let t=0;t<tabs.length;t++)tabs[t].addEventListener("click",function(){let t=this.getAttribute("data-id"),i=document.getElementById(t);document.querySelector(".it-block__tab.active").classList.remove("active"),document.querySelector(".it-block__plane.active").classList.remove("active"),this.classList.add("active"),i.classList.add("active")});'}],
  });
}
</script>

<template>
  <div class="intro">
    <div class="text-wrap">
      <h2 class="header">Добро пожаловать на WB Tech</h2>
      <p class="text">Demo for Level 0</p>
    </div>
    <div class="banner">
      <div class="canvas-wrap">
        <canvas id="bubble" width="1440" height="422"/>
      </div>
    </div>
  </div>
</template>


<style scoped>
.intro {
  position: relative;
  padding: 25px 120px 35px;
  margin-top: 20px;
}

.text-wrap {
  position: relative;
  z-index: 2;
  margin: 0;
}

.header {
  font-size: 50px;
  line-height: 75px;
  max-width: 700px;
  font-weight: bold;
  color: #ffffff;
}
.text {
  font-size: 18px;
  line-height: 24px;
  max-width: 375px;
  color: #ffffff;
}
.banner {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

.canvas-wrap {
  overflow: hidden;
  position: relative;
  height: 100%;
  border-radius: 50px;
  background: linear-gradient(150deg, #DD37D0, #8C1FF2);
}

.canvas-wrap canvas {
  width: 100%;
  height: 100%;
}
</style>