(window.webpackJsonp=window.webpackJsonp||[]).push([[10],{112:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return c})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return s}));var n=r(1),o=r(6),a=(r(0),r(176)),c={title:"UI"},i={id:"reference/ui",title:"UI",description:"## Base URL",source:"@site/docs/reference/ui.md",permalink:"/docs/reference/ui",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/ui.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1588472011,sidebar:"docs",previous:{title:"String",permalink:"/docs/reference/opspec/types/string"},next:{title:"global-options",permalink:"/docs/reference/cli/global-options"}},p=[{value:"Base URL",id:"base-url",children:[]},{value:"Query Parameters",id:"query-parameters",children:[{value:"<code>mount</code>",id:"mount",children:[]}]}],l={rightToc:p},u="wrapper";function s(e){var t=e.components,r=Object(o.a)(e,["components"]);return Object(a.b)(u,Object(n.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(a.b)("h2",{id:"base-url"},"Base URL"),Object(a.b)("p",null,Object(a.b)("a",Object(n.a)({parentName:"p"},{href:"http://localhost:42224"}),"http://localhost:42224"),"."),Object(a.b)("h2",{id:"query-parameters"},"Query Parameters"),Object(a.b)("h3",{id:"mount"},Object(a.b)("inlineCode",{parentName:"h3"},"mount")),Object(a.b)("p",null,"URL encoded reference to mount (either ",Object(a.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(a.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(a.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(a.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")"),Object(a.b)("h4",{id:"example-mount-githubcomopspec-pkgs_opcreate331"},"Example mount github.com/opspec-pkgs/_.op.create#3.3.1"),Object(a.b)("ol",null,Object(a.b)("li",{parentName:"ol"},"Navigate to ",Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"http://localhost:42224/?mount=github.com%2Fopspec-pkgs%2F_.op.create%233.3.1"}),"http://localhost:42224/?mount=github.com%2Fopspec-pkgs%2F_.op.create%233.3.1"))))}s.isMDXComponent=!0},176:function(e,t,r){"use strict";r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return f}));var n=r(0),o=r.n(n);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=o.a.createContext({}),u=function(e){var t=o.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},s=function(e){var t=u(e.components);return(o.a.createElement(l.Provider,{value:t},e.children))},b="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return o.a.createElement(o.a.Fragment,{},t)}},d=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,a=e.originalType,c=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),s=u(r),b=n,d=s["".concat(c,".").concat(b)]||s[b]||m[b]||a;return r?o.a.createElement(d,i({ref:t},l,{components:r})):o.a.createElement(d,i({ref:t},l))}));function f(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var a=r.length,c=new Array(a);c[0]=d;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[b]="string"==typeof e?e:n,c[1]=i;for(var l=2;l<a;l++)c[l]=r[l];return o.a.createElement.apply(null,c)}return o.a.createElement.apply(null,r)}d.displayName="MDXCreateElement"}}]);