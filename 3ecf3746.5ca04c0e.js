(window.webpackJsonp=window.webpackJsonp||[]).push([[17],{110:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return i})),r.d(t,"metadata",(function(){return b})),r.d(t,"rightToc",(function(){return o})),r.d(t,"default",(function(){return d}));var a=r(1),n=r(6),c=(r(0),r(147)),i={sidebar_label:"Index",title:"Container Call [object]"},b={id:"opspec/reference/structure/op-directory/op/call/container/index",title:"Container Call [object]",description:"An object defining a container call.",source:"@site/docs/opspec/reference/structure/op-directory/op/call/container/index.md",permalink:"/docs/opspec/reference/structure/op-directory/op/call/container/index",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/call/container/index.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1585210706,sidebar_label:"Index",sidebar:"docs",previous:{title:"Call [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/index"},next:{title:"Image [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/container/image"}},o=[{value:"Properties",id:"properties",children:[{value:"image",id:"image",children:[]},{value:"cmd",id:"cmd",children:[]},{value:"dirs",id:"dirs",children:[]},{value:"envVars",id:"envvars",children:[]},{value:"files",id:"files",children:[]},{value:"name",id:"name",children:[]},{value:"ports",id:"ports",children:[]},{value:"sockets",id:"sockets",children:[]},{value:"workDir",id:"workdir",children:[]}]}],l={rightToc:o},p="wrapper";function d(e){var t=e.components,r=Object(n.a)(e,["components"]);return Object(c.b)(p,Object(a.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object defining a container call."),Object(c.b)("h2",{id:"properties"},"Properties"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"must have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#image"}),"image")))),Object(c.b)("li",{parentName:"ul"},"may have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#cmd"}),"cmd")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#dirs"}),"dirs")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#envvars"}),"envVars")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#files"}),"files")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#name"}),"name")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#ports"}),"ports")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#sockets"}),"sockets")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(a.a)({parentName:"li"},{href:"#workdir"}),"workDir"))))),Object(c.b)("h3",{id:"image"},"image"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/container/image"}),"image [object]")," defining the container image run by the call."),Object(c.b)("h3",{id:"cmd"},"cmd"),Object(c.b)("p",null,"An array of ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/types/string#initialization"}),"string initializers")," defining the path (from ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"#workdir"}),"workDir"),") of the binary to call and it's arguments."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"defining cmd overrides any entrypoint and/or cmd defined by the image")),Object(c.b)("h3",{id:"dirs"},"dirs"),Object(c.b)("p",null,"An object for which each key is an absolute path in the container and each value is one of:"),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"value"),Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"meaning"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Mount dir embedded in op w/ same path (equivalent to ",Object(c.b)("inlineCode",{parentName:"td"},"$(/absolute/path)"),")")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(a.a)({parentName:"td"},{href:"/docs/opspec/reference/types/dir"}),"dir")," ",Object(c.b)("a",Object(a.a)({parentName:"td"},{href:"/docs/opspec/reference/structure/op-directory/op/variable-reference"}),"variable-reference [string]")),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Mount dir")))),Object(c.b)("h3",{id:"envvars"},"envVars"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/types/object#initialization"}),"object initializer")," or ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/variable-reference"}),"variable-reference [string]"),", whos properties represent the name and value of an environment variable to be set in the container."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"upon evaluation, the key and value of each property will be coerced to a string.")),Object(c.b)("h3",{id:"files"},"files"),Object(c.b)("p",null,"An object for which each key is an absolute path in the container and each value is one of:"),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"value"),Object(c.b)("th",Object(a.a)({parentName:"tr"},{align:null}),"meaning"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Mount file embedded in op w/ same path (equivalent to ",Object(c.b)("inlineCode",{parentName:"td"},"$(/absolute/path)"),")")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(a.a)({parentName:"td"},{href:"/docs/opspec/reference/types/file"}),"file")," ",Object(c.b)("a",Object(a.a)({parentName:"td"},{href:"/docs/opspec/reference/structure/op-directory/op/variable-reference"}),"variable-reference [string]")),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Mount file")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(a.a)({parentName:"td"},{href:"/docs/opspec/reference/types/file#initialization"}),"file initializer")),Object(c.b)("td",Object(a.a)({parentName:"tr"},{align:null}),"Evaluate and mount")))),Object(c.b)("h3",{id:"name"},"name"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/types/string#initialization"}),"string initializer")," defining a name by which the container can be resolved on the opctl network."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"if multiple containers are given the same name, network requests will be distributed (load balanced) across them. ")),Object(c.b)("h3",{id:"ports"},"ports"),Object(c.b)("p",null,"An object defining container ports exposed on the opctl host where:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"each key is a container port or range of ports (optionally including protocol) matching ",Object(c.b)("inlineCode",{parentName:"li"},"[0-9]+(-[0-9]+)?(tcp|udp)")),Object(c.b)("li",{parentName:"ul"},"each value is a corresponding opctl host port or range of ports matching ",Object(c.b)("inlineCode",{parentName:"li"},"[0-9]+(-[0-9]+)?"))),Object(c.b)("h3",{id:"sockets"},"sockets"),Object(c.b)("p",null,"An object for which each key is an absolute path in the container and and each value is a ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/types/socket"}),"socket")," ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/variable-reference"}),"variable-reference [string]")," to mount. "),Object(c.b)("h3",{id:"workdir"},"workDir"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"/docs/opspec/reference/types/string#initialization"}),"string initializer")," defining absolute path from which ",Object(c.b)("a",Object(a.a)({parentName:"p"},{href:"#cmd"}),"cmd")," will be executed."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"defining workDir overrides any defined by the image")))}d.isMDXComponent=!0},147:function(e,t,r){"use strict";r.d(t,"a",(function(){return d})),r.d(t,"b",(function(){return O}));var a=r(0),n=r.n(a);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,a)}return r}function b(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function o(e,t){if(null==e)return{};var r,a,n=function(e,t){if(null==e)return{};var r,a,n={},c=Object.keys(e);for(a=0;a<c.length;a++)r=c[a],t.indexOf(r)>=0||(n[r]=e[r]);return n}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(a=0;a<c.length;a++)r=c[a],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(n[r]=e[r])}return n}var l=n.a.createContext({}),p=function(e){var t=n.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):b({},t,{},e)),r},d=function(e){var t=p(e.components);return n.a.createElement(l.Provider,{value:t},e.children)},s="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.a.createElement(n.a.Fragment,{},t)}},j=Object(a.forwardRef)((function(e,t){var r=e.components,a=e.mdxType,c=e.originalType,i=e.parentName,l=o(e,["components","mdxType","originalType","parentName"]),d=p(r),s=a,j=d["".concat(i,".").concat(s)]||d[s]||u[s]||c;return r?n.a.createElement(j,b({ref:t},l,{components:r})):n.a.createElement(j,b({ref:t},l))}));function O(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var c=r.length,i=new Array(c);i[0]=j;var b={};for(var o in t)hasOwnProperty.call(t,o)&&(b[o]=t[o]);b.originalType=e,b[s]="string"==typeof e?e:a,i[1]=b;for(var l=2;l<c;l++)i[l]=r[l];return n.a.createElement.apply(null,i)}return n.a.createElement.apply(null,r)}j.displayName="MDXCreateElement"}}]);