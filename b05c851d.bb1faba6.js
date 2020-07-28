(window.webpackJsonp=window.webpackJsonp||[]).push([[39],{178:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return o})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return b})),r.d(t,"default",(function(){return u}));var a=r(1),n=r(9),c=(r(0),r(197)),o={title:"Loop Vars [object]"},i={id:"reference/opspec/op-directory/op/call/loop-vars",title:"Loop Vars [object]",description:"An object naming variables to assign each iterations info to.",source:"@site/docs/reference/opspec/op-directory/op/call/loop-vars.md",permalink:"/docs/reference/opspec/op-directory/op/call/loop-vars",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/op-directory/op/call/loop-vars.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1587672399,sidebar:"docs",previous:{title:"Image [object]",permalink:"/docs/reference/opspec/op-directory/op/call/container/image"},next:{title:"Op Call [object]",permalink:"/docs/reference/opspec/op-directory/op/call/op"}},b=[{value:"Properties",id:"properties",children:[{value:"index",id:"index",children:[]},{value:"key",id:"key",children:[]},{value:"value",id:"value",children:[]}]}],l={rightToc:b},p="wrapper";function u(e){var t=e.components,r=Object(n.a)(e,["components"]);return Object(c.b)(p,Object(a.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object naming variables to assign each iterations info to."),Object(c.b)("h2",{id:"properties"},"Properties"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"may have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#index"}),"index")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#key"}),"key")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#value"}),"value"))))),Object(c.b)("h3",{id:"index"},"index"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/opspec/op-directory/op/identifier"}),"identifier [string]")," naming a variable to bind each iterations index to."),Object(c.b)("h3",{id:"key"},"key"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/opspec/op-directory/op/identifier"}),"identifier [string]")," naming a variable to bind each iterations key to."),Object(c.b)("p",null,"Behavior varies based on the range value:  "),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"range"),Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"behavior"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Variable not set")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"array"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Variable set to current item index")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"object"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Variable set to current property name")))),Object(c.b)("h3",{id:"value"},"value"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/reference/opspec/op-directory/op/identifier"}),"identifier [string]")," naming a variable to bind each iterations value to."),Object(c.b)("p",null,"Behavior varies based on the range value:  "),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"range value"),Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"behavior"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Variable not set")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"array"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Variable set to current item")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"object"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Variable set to current property value")))))}u.isMDXComponent=!0},197:function(e,t,r){"use strict";r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return s}));var a=r(0),n=r.n(a);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,a)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function b(e,t){if(null==e)return{};var r,a,n=function(e,t){if(null==e)return{};var r,a,n={},c=Object.keys(e);for(a=0;a<c.length;a++)r=c[a],t.indexOf(r)>=0||(n[r]=e[r]);return n}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(a=0;a<c.length;a++)r=c[a],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(n[r]=e[r])}return n}var l=n.a.createContext({}),p=function(e){var t=n.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},u=function(e){var t=p(e.components);return(n.a.createElement(l.Provider,{value:t},e.children))},d="mdxType",O={inlineCode:"code",wrapper:function(e){var t=e.children;return n.a.createElement(n.a.Fragment,{},t)}},j=Object(a.forwardRef)((function(e,t){var r=e.components,a=e.mdxType,c=e.originalType,o=e.parentName,l=b(e,["components","mdxType","originalType","parentName"]),u=p(r),d=a,j=u["".concat(o,".").concat(d)]||u[d]||O[d]||c;return r?n.a.createElement(j,i({ref:t},l,{components:r})):n.a.createElement(j,i({ref:t},l))}));function s(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var c=r.length,o=new Array(c);o[0]=j;var i={};for(var b in t)hasOwnProperty.call(t,b)&&(i[b]=t[b]);i.originalType=e,i[d]="string"==typeof e?e:a,o[1]=i;for(var l=2;l<c;l++)o[l]=r[l];return n.a.createElement.apply(null,o)}return n.a.createElement.apply(null,r)}j.displayName="MDXCreateElement"}}]);