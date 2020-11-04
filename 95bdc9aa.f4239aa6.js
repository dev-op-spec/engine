(window.webpackJsonp=window.webpackJsonp||[]).push([[37],{139:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return p})),n.d(t,"metadata",(function(){return c})),n.d(t,"rightToc",(function(){return i})),n.d(t,"default",(function(){return s}));var o=n(1),r=n(6),a=(n(0),n(169)),p={},c={id:"reference/cli/op",title:"op",description:"## `opctl op create [OPTIONS] NAME`",source:"@site/docs/reference/cli/op.md",permalink:"/docs/reference/cli/op",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/cli/op.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1604375618,sidebar:"docs",previous:{title:"node",permalink:"/docs/reference/cli/node"},next:{title:"run",permalink:"/docs/reference/cli/run"}},i=[{value:"<code>opctl op create [OPTIONS] NAME</code>",id:"opctl-op-create-options-name",children:[{value:"Arguments",id:"arguments",children:[]},{value:"Options",id:"options",children:[]},{value:"Examples",id:"examples",children:[]}]},{value:"<code>opctl op install [OPTIONS] OP_REF</code>",id:"opctl-op-install-options-op_ref",children:[{value:"Arguments",id:"arguments-1",children:[]},{value:"Options",id:"options-1",children:[]},{value:"Examples",id:"examples-1",children:[]},{value:"Notes",id:"notes",children:[]}]},{value:"<code>opctl op kill OP_ID</code>",id:"opctl-op-kill-op_id",children:[{value:"Arguments",id:"arguments-2",children:[]}]},{value:"<code>opctl op validate OP_REF</code>",id:"opctl-op-validate-op_ref",children:[{value:"Arguments",id:"arguments-3",children:[]},{value:"Examples",id:"examples-2",children:[]},{value:"Notes",id:"notes-1",children:[]}]}],l={rightToc:i},b="wrapper";function s(e){var t=e.components,n=Object(r.a)(e,["components"]);return Object(a.b)(b,Object(o.a)({},l,n,{components:t,mdxType:"MDXLayout"}),Object(a.b)("h2",{id:"opctl-op-create-options-name"},Object(a.b)("inlineCode",{parentName:"h2"},"opctl op create [OPTIONS] NAME")),Object(a.b)("p",null,"Creates an op"),Object(a.b)("h3",{id:"arguments"},"Arguments"),Object(a.b)("h4",{id:"name"},Object(a.b)("inlineCode",{parentName:"h4"},"NAME")),Object(a.b)("p",null,"Name of the op"),Object(a.b)("h3",{id:"options"},"Options"),Object(a.b)("h4",{id:"-d-or---description"},Object(a.b)("inlineCode",{parentName:"h4"},"-d")," or ",Object(a.b)("inlineCode",{parentName:"h4"},"--description")),Object(a.b)("p",null,"Description of the op"),Object(a.b)("h4",{id:"--path-default-opspec"},Object(a.b)("inlineCode",{parentName:"h4"},"--path")," ",Object(a.b)("em",{parentName:"h4"},"default: ",Object(a.b)("inlineCode",{parentName:"em"},".opspec"))),Object(a.b)("p",null,"Path to create the op at"),Object(a.b)("h3",{id:"examples"},"Examples"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-sh"}),'opctl op create -d "my awesome op description" --path some/path my-awesome-op-name\n')),Object(a.b)("h2",{id:"opctl-op-install-options-op_ref"},Object(a.b)("inlineCode",{parentName:"h2"},"opctl op install [OPTIONS] OP_REF")),Object(a.b)("p",null,"Installs an op"),Object(a.b)("h3",{id:"arguments-1"},"Arguments"),Object(a.b)("h4",{id:"op_ref"},Object(a.b)("inlineCode",{parentName:"h4"},"OP_REF")),Object(a.b)("p",null,"Op reference (",Object(a.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(a.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")"),Object(a.b)("h3",{id:"options-1"},"Options"),Object(a.b)("h4",{id:"--path-default-opspecop_ref"},Object(a.b)("inlineCode",{parentName:"h4"},"--path")," ",Object(a.b)("em",{parentName:"h4"},"default: ",Object(a.b)("inlineCode",{parentName:"em"},".opspec/OP_REF"))),Object(a.b)("p",null,"Path to install the op at"),Object(a.b)("h4",{id:"-u-or---username"},Object(a.b)("inlineCode",{parentName:"h4"},"-u")," or ",Object(a.b)("inlineCode",{parentName:"h4"},"--username")),Object(a.b)("p",null,"Username used to auth w/ the op source"),Object(a.b)("h4",{id:"-p-or---password"},Object(a.b)("inlineCode",{parentName:"h4"},"-p")," or ",Object(a.b)("inlineCode",{parentName:"h4"},"--password")),Object(a.b)("p",null,"Password used to auth w/ the op source"),Object(a.b)("h3",{id:"examples-1"},"Examples"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-sh"}),"opctl op install -u someUser -p somePass host/path/repo#tag\n")),Object(a.b)("h3",{id:"notes"},"Notes"),Object(a.b)("h4",{id:"op-source-usernamepassword-prompt"},"op source username/password prompt"),Object(a.b)("p",null,"If auth w/ the op source fails the cli will (re)prompt for username &\npassword."),Object(a.b)("blockquote",null,Object(a.b)("p",{parentName:"blockquote"},"in non-interactive terminals, the cli will note that it can't prompt\nand exit with a non zero exit code.")),Object(a.b)("h2",{id:"opctl-op-kill-op_id"},Object(a.b)("inlineCode",{parentName:"h2"},"opctl op kill OP_ID")),Object(a.b)("p",null,"Kill an op. "),Object(a.b)("h3",{id:"arguments-2"},"Arguments"),Object(a.b)("h4",{id:"op_id"},Object(a.b)("inlineCode",{parentName:"h4"},"OP_ID")),Object(a.b)("p",null,"Id of the op to kill"),Object(a.b)("h2",{id:"opctl-op-validate-op_ref"},Object(a.b)("inlineCode",{parentName:"h2"},"opctl op validate OP_REF")),Object(a.b)("p",null,"Validates an op according to:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"existence of ",Object(a.b)("inlineCode",{parentName:"li"},"op.yml")),Object(a.b)("li",{parentName:"ul"},"validity of ",Object(a.b)("inlineCode",{parentName:"li"},"op.yml")," (per\n",Object(a.b)("a",Object(o.a)({parentName:"li"},{href:"https://opctl.io/0.1.6/op.yml.schema.json"}),"schema"),")")),Object(a.b)("h3",{id:"arguments-3"},"Arguments"),Object(a.b)("h4",{id:"op_ref-1"},Object(a.b)("inlineCode",{parentName:"h4"},"OP_REF")),Object(a.b)("p",null,"Op reference (either ",Object(a.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(a.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(a.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(a.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")."),Object(a.b)("h3",{id:"examples-2"},"Examples"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-sh"}),"opctl op validate myop\n")),Object(a.b)("h3",{id:"notes-1"},"Notes"),Object(a.b)("h4",{id:"op-source-usernamepassword-prompt-1"},"op source username/password prompt"),Object(a.b)("p",null,"If auth w/ the op source fails the cli will (re)prompt for username & password."),Object(a.b)("blockquote",null,Object(a.b)("p",{parentName:"blockquote"},"in non-interactive terminals, the cli will note that it can't prompt and exit with a non zero exit code.")))}s.isMDXComponent=!0},169:function(e,t,n){"use strict";n.d(t,"a",(function(){return s})),n.d(t,"b",(function(){return O}));var o=n(0),r=n.n(o);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function p(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function c(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?p(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):p(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},a=Object.keys(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var l=r.a.createContext({}),b=function(e){var t=r.a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):c({},t,{},e)),n},s=function(e){var t=b(e.components);return(r.a.createElement(l.Provider,{value:t},e.children))},d="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.a.createElement(r.a.Fragment,{},t)}},m=Object(o.forwardRef)((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,p=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),s=b(n),d=o,m=s["".concat(p,".").concat(d)]||s[d]||u[d]||a;return n?r.a.createElement(m,c({ref:t},l,{components:n})):r.a.createElement(m,c({ref:t},l))}));function O(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,p=new Array(a);p[0]=m;var c={};for(var i in t)hasOwnProperty.call(t,i)&&(c[i]=t[i]);c.originalType=e,c[d]="string"==typeof e?e:o,p[1]=c;for(var l=2;l<a;l++)p[l]=n[l];return r.a.createElement.apply(null,p)}return r.a.createElement.apply(null,n)}m.displayName="MDXCreateElement"}}]);