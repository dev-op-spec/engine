(window.webpackJsonp=window.webpackJsonp||[]).push([[37],{130:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return i})),r.d(t,"metadata",(function(){return c})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return u}));var n=r(1),a=r(6),o=(r(0),r(147)),i={title:"Travis",sidebar_label:"Travis"},c={id:"setup/travis",title:"Travis",description:'[travis ci](https://travis-ci.org/) looks for a `travis.yml` file at the root of each repo to identify ci "stages".',source:"@site/docs/setup/travis.md",permalink:"/docs/setup/travis",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/setup/travis.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1587014832,sidebar_label:"Travis",sidebar:"docs",previous:{title:"Kubernetes",permalink:"/docs/setup/kubernetes"},next:{title:"Hello World",permalink:"/docs/zero-to-hero/hello-world"}},l=[{value:"Examples",id:"examples",children:[]}],s={rightToc:l},p="wrapper";function u(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(o.b)(p,Object(n.a)({},s,r,{components:t,mdxType:"MDXLayout"}),Object(o.b)("p",null,Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"https://travis-ci.org/"}),"travis ci")," looks for a ",Object(o.b)("inlineCode",{parentName:"p"},"travis.yml"),' file at the root of each repo to identify ci "stages".'),Object(o.b)("p",null,"Their hosted agents support starting the ci process alongside a docker daemon so running opctl is\na matter of defining your ",Object(o.b)("inlineCode",{parentName:"p"},"travis.yml")," as follows:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"adding opctl installation to the ",Object(o.b)("inlineCode",{parentName:"li"},"before_script")," array"),Object(o.b)("li",{parentName:"ul"},"adding docker to the ",Object(o.b)("inlineCode",{parentName:"li"},"services")," array"),Object(o.b)("li",{parentName:"ul"},"adding ",Object(o.b)("inlineCode",{parentName:"li"},"script")," array entries with your calls to opctl")),Object(o.b)("h3",{id:"examples"},"Examples"),Object(o.b)("p",null,"travis.yml"),Object(o.b)("pre",null,Object(o.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"language: generic\nsudo: required\nbefore_script:\n- curl -L https://github.com/opctl/opctl/releases/download/0.1.31/opctl0.1.31.linux.tgz | sudo tar -xzv -C /usr/local/bin\nservices:\n- docker\nscript:\n# passes an arg `gitBranch` to the op from a travis-ci variable\n- opctl run -a gitBranch=$TRAVIS_BRANCH build\n")))}u.isMDXComponent=!0},147:function(e,t,r){"use strict";r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return m}));var n=r(0),a=r.n(n);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function c(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=a.a.createContext({}),p=function(e){var t=a.a.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):c({},t,{},e)),r},u=function(e){var t=p(e.components);return a.a.createElement(s.Provider,{value:t},e.children)},b="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,o=e.originalType,i=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=p(r),b=n,f=u["".concat(i,".").concat(b)]||u[b]||d[b]||o;return r?a.a.createElement(f,c({ref:t},s,{components:r})):a.a.createElement(f,c({ref:t},s))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=r.length,i=new Array(o);i[0]=f;var c={};for(var l in t)hasOwnProperty.call(t,l)&&(c[l]=t[l]);c.originalType=e,c[b]="string"==typeof e?e:n,i[1]=c;for(var s=2;s<o;s++)i[s]=r[s];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);