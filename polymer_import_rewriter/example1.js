import {  html, css } from 'lit-element';
import '@polymer/iron-flex-layout/iron-flex-layout.js';
import { View } from '../../view/view.js';
import '../../view/view-search/view-search.js';
import '../dataset/data-source.js';

import '@material/mwc-tab-indicator';
import { addHasRemoveClass, BaseElement } from '@material/mwc-base/base-element.js';
import { ripple } from '@material/mwc-ripple/ripple-directive';
import MDCTabFoundation from '@material/tab/foundation';
import { html, property, query } from 'lit-element';
import { classMap } from 'lit-html/directives/class-map';
import { style } from './mwc-tab-css';
import { __decorate } from "tslib";

/**
 * `view-tree`
 * 
 *
 * @customElement
 * @polymer
 * @demo demo/index.html
 */
class AppViewControlPanel extends View {
    static get styles() {
        return [
        super.styles,
        css`
                :host {
                display: block;
                -webkit-box-flex: 0;
                -webkit-flex: 0 0 auto;
                -ms-flex: 0 0 auto;
                flex: 0 0 auto;
                user-select: none;
                padding-top: 10px;
                padding-right: 16px;
                padding-bottom: 10px;
                padding-left: 16px;
                background-color: white;
                border-bottom: 1px solid #a8a8a8;
            }
            .container-fluid {
                width: 100%;
            }

            .container-fluid>div {
                width: 50%;
            }

            .oe-button-column {
                height: 30px;
            }

            :host(.dropdown-menu) {
                max-height: calc(100vh - 125px);
                overflow: auto;
                overflow-x: hidden;
            }

                :host(.dropdown-menu) li {
                position: relative;
            }

                :host(.dropdown-menu) li a {
                padding: 3px 25px;
            }

                :host(.dropdown-menu).oe_searchview_custom_public a:after,
                :host(.dropdown-menu) .oe-share-filter:after {
                font-family: FontAwesome;
                content: "";
                color: #666;
                margin-left: 3px;
            }

                :host(.selected) {
                display: block;
            }

                :host(.selected) a {
                font-weight: bold;
            }

                :host(.selected) a:before {
                font-family: FontAwesome;
                position: absolute;
                left: 6px;
                top: 3px;
                content: "";
            }

                :host(.oe_tag) {
                -moz-border-radius: 0px;
                -webkit-border-radius: 0px;
                border-radius: 0px;
            }

                :host(.oe-view-title) {
                font-size: 18px;
                padding-left: 0;
                margin: 0;
                background-color: #f0eeee;
            }

                :host(.oe-view-title) li {
                -moz-user-select: initial;
                -webkit-user-select: initial;
                user-select: initial;
            }
            

            :host(.cp-search-view) {
                padding-top: 5px;
            }

                :host(.cp-modes) button {
                width: 34px;
            }

                :host(.cp-buttons) {
                display: inline-block;
            }

                :host().cp-sidebar) {
                display: inline-block;
                float: right;
            }

                :host(.cp-sidebar) .o_form_binary_form {
                cursor: pointer;
            }

                :host(.cp-sidebar) .o_form_binary_form span {
                padding: 3px 20px;
            }

                :host(.cp-sidebar) .o_form_binary_form input.o_form_input_file {
                width: 100%;
            }

                :host(.cp-sidebar) .o_form_binary_form:hover {
                background-color: #f5f5f5;
            }

                :host(.cp-sidebar) .oe_file_attachment {
                padding: 3px 20px;
                display: inline-block;
            }

                :host(.cp-sidebar) .oe_sidebar_delete_item {
                padding: 0;
                position: absolute;
                right: 10px;
            }

                :host(.cp-sidebar) .dropdown-menu li a {
                width: 100%;
            }

                :host(.form_buttons) {
                padding: 0;
            }

                :host(.o_form_buttons_view)>button {
                float: left;
            }

                :host(.o_form_buttons_view)>button:last-child {
                float: right;
                margin-left: 4px;
            }

                :host(.pager-buttons) {
                min-height: 30px;
            }

                :host * {
                -webkit-box-sizing: border-box;
                -moz-box-sizing: border-box;
                box-sizing: border-box;
            }

                :host(.search-on) {
                left: 0;
                background: inherit;
                z-index: 1001;
            }

            ::slotted(iron-icon) {
                margin-right: 15px;
                cursor: pointer;
            }

            #app-control-panel {
                position: relative;
            }

            #app-control-panel iron-icon {
                margin-right: 0;
            }

            #search[show] {
                position: absolute;
                width: 100%;
                height: 100%;
                left: 0;
                top: 0;
                padding: 0 16px;
                background: #fff;
            }

            #search input {
                display: none;
                font-family: var(--primary-font-family);
                font-size: 15px;
                width: 100%;
                padding: 10px;
                border: 0;
                border-radius: 2px;
                -webkit-appearance: none;
            }

            #search[show] input {
                display: block;
            }

            #search input:focus {
                outline: 0;
            }

                :host(.container-fluid) {
                padding-top: 10px;
                padding-right: 16px;
                padding-bottom: 10px;
                padding-left: 16px;
                background-color: white;
                border-bottom: 1px solid #a8a8a8;
                display: -ms-flexbox;
                display: -moz-box;
                display: -webkit-box;
                display: -webkit-flex;
                display: flex;
                -ms-flex-flow: row wrap;
                -moz-flex-flow: row wrap;
                -webkit-flex-flow: row wrap;
                flex-flow: row wrap;
            }



            .row:first-child {
                padding-top: 3px;
                padding-bottom: 3px;
            }

            .row:last-child {
                padding-bottom: 10px;
            }

                :host(.col-md-6) {
                width: 50%;
            }
            

            .btn {
                display: inline-block;
                margin-bottom: 0;
                font-weight: 400;
                text-align: center;
                vertical-align: middle;
                touch-action: manipulation;
                cursor: pointer;
                background-image: none;
                border: 1px solid transparent;
                white-space: nowrap;
                padding: 6px 12px;
                font-size: 14px;
                line-height: 1.42857143;
                border-radius: 4px;
                -webkit-user-select: none;
                -moz-user-select: none;
                -ms-user-select: none;
                user-select: none;
            }
            .btn-sm,
            .btn-group-sm>.btn {
                padding: 5px 10px;
                font-size: 12px;
                line-height: 1.5;
                border-radius: 3px;
            }

                .btn-primary {
                color: white;
                background-color: #21b799;
                border-color: #21b799;
            }

                .btn-primary:hover {
                color: white;
                background-color: #198c75;
                border-color: #18836e;
            }

            .btn,
        .btn-sm,
        .btn-group>.btn,
        .btn-group-sm>.btn {
            border-radius: 0px;
            border: none;
        }

        .btn-default {
            color: #333;
            background-color: #fff;
            border-color: #ccc;
        }

        .btn-default.active,
        .btn-default:active,
        .open>.dropdown-toggle.btn-default {
            color: #333;
            background-color: #e6e6e6;
            border-color: #adadad;
        }

        .btn-group>.btn {
            position: relative;
            float: left;
        }
        :host (.cp-tree) {
        display: -ms-flexbox;
        display: -moz-box;
        display: -webkit-box;
        display: -webkit-flex;
        display: flex;
        -moz-justify-content: space-between;
        -webkit-justify-content: space-between;
        justify-content: space-between;
    }

        :host (.cp-sidebar) {
        padding-right: 10px;
    }

        :host (.cp-pager) {
        margin: auto 0 auto auto;
    }
    

    ::slotted(.cp-breadcrumb),
    .cp-breadcrumb {
        padding: 0 0;
        margin-bottom: 18px;
        list-style: none;
    }

    ::slotted(.cp-breadcrumb)>li+li:before {
        /*content: "/\00a0";*/
        padding: 0 0px 0 5px;
        color: #777777;
    }
`];
    }

    render() {
        return html`
                <div id="container-fluid" class="container-fluid horizontal layout wrap">
                    <!-- session one-->
                    <div class="cp-one col-md-6 cp-title ">
                        <ol id="view-title" class="cp-view-title cp-breadcrumb horizontal layout" restamp></ol>
                    </div>
                    <!-- session two-->
                    <div class="cp-two cp-search-view col-md-6"></div>
                    <!-- session tree-->
                    <div class="cp-tree col-md-6 button-column ">
                        <div class="cp-buttons"></div>
                        <div class="cp-sidebar"></div>
                    </div>
                    <!-- session four-->
                    <div class="cp-four col-md-6 horizontal layout center ">
                        <!-- 搜索选项-->
                        <div class="cp-search-options"> </div>

                        <!-- 页数显示-->
                        <div class="cp-pager"> </div>

                        <!-- mode 选择-->
                        <div class="cp-modes btn-group btn-group-sm right-toolbar">
                            <dom-repeat items="[[modes]]" as="mode">
                                <!-- 根据不同View添加不同图标-->
                                <paper-icon-button icon="view-[[mode]]-icon:[[mode]]" mode-type="[[mode]]" title="[[mode]]" on-tap="onModeBtnClick"></paper-icon-button>
                            </dom-repeat>
                        </div>
                    </div>
                </div>
                <!--Search View-->
                <!--template is="dom-if" if="[[action]]" restamp>
                <data-source id="dataset" action="/dataset/call_kw/[[action.res_model]]/fields_view_get" params='{"view_id":[[action.search_view_id]],"model":"[[action.res_model]]","view_type":"search"}' data="{{view}}" index=0 active></data-source>
            </template-->
                <slot></slot>
            `;
    }

    static get properties() {
        return {
            action: {
                type: Object,
                observer: 'onActionChanged'
            },

            fields: { // # 该Action可用的所有字段
                type: Object,
            },

            // 视图集
            views: {
                type: Object,
                observer: 'onViewsChanged'
            },

            // 视图
            view: {
                type: Object,
                observer: 'onViewChanged'
            },

            modes: {
                type: Object,
            },

            showingSearch: {
                type: Boolean,
                value: false
            },
        };
    }

    static get observers() {
        return [
            //  'updateSearchDisplay(showingSearch)',
            'onViewsLoaded(views)',
        ];
    }
    /*
    listeners: {
     keyup: 'onHotkeys'
 },
    */

   firstUpdated() {
        //super();

        // # 初始化控制元素
        var a = this.app;
        this.controlElements = {};
        var shadow = this.shadowRoot

        // # 主控制板块
        this.controlElements.cp_one = shadow.querySelector('div.cp-one');
        this.controlElements.cp_two = shadow.querySelector('div.cp-two');
        this.controlElements.cp_tree = shadow.querySelector('div.cp-tree');
        this.controlElements.cp_four = shadow.querySelector('div.cp-four');
        // # 二级板块
        this.controlElements.breadcrumb = shadow.querySelector('ol.cp-breadcrumb');
        this.controlElements.searchview = shadow.querySelector('div.cp-search-view');
        this.controlElements.searchviewButtons = shadow.querySelector('div.cp-search-options');
        this.controlElements.buttons = shadow.querySelector('div.cp-buttons');
        this.controlElements.sidebar = shadow.querySelector('div.cp-sidebar');
        this.controlElements.pager = shadow.querySelector('div.cp-pager');
        this.controlElements.switchButtons = shadow.querySelector('div.cp-modes');

        //this.listen(this, "search_data", "doSearch"); // #监听SearchView查询事件
    }

    doSearch(event, ctx) {
        if (!this.fields) {
            return
        }

        var self = this;
        //  var controller = this.active_view.controller; // the correct view must be loaded here
        var action_context = this.action.context || {};
        var view_context = this.get_context();
        pyeval.eval_domains_and_contexts({
            domains: [this.action.domain || []].concat(ctx.domains || []),
            contexts: [action_context, view_context].concat(ctx.contexts || []),
            group_by_seq: ctx.groupbys || []
        }).then(function (results) {
            if (results.error) {
                // self.active_search.resolve();
                // throw new Error(
                //         _.str.sprintf(_t("Failed to evaluate search criterions")+": \n%s",
                //                       JSON.stringify(results.error)));
            }
            //self.dataset._model = new Model(
            //   self.dataset.model, results.context, results.domain);
            var groupby = results.group_by.length ?
                results.group_by :
                action_context.group_by;
            if (vectors.utils.isString(groupby)) {
                groupby = [groupby];
            }
            //if (!controller.grouped && !vectors.utils.isEmpty(groupby)) {
            //    self.dataset.set_sort([]);
            // }

            results.group_by = groupby;
            results.views = self.views.fields_views; // # 传递Views视图列表给ViewManager
            results.fields = self.fields; // # 传递Views字段列表给ViewManager
            results.model = self.action.res_model;
            results.limit = self.action.limit;
            results.offset = '';
            results.sort = '';
            // # 触发查询事件
            //self.fire("do-search", results);
            this.dispatchEvent(new CustomEvent('search', results));

            //$.when(controller.do_search(results.domain, results.context, groupby || [])).then(function () {
            //    self.active_search.resolve();
            // });
        });
    }

    onViewsLoaded(views) {
        this.fields = views.fields;
        if ("search" in views.fields_views) {
            this.view = views.fields_views["search"];
        }
    }

    onViewChanged(view) {
        var lSearchView = this.querySelector('div.cp-search-view'); //
        lSearchView.innerHTML = view.arch;


        // 展示View
        var lVs = Polymer.dom(lSearchView).firstChild;
        lVs.viewMgr = this;
        lVs.show(this);
        //lVs=  this.queryEffectiveChildren('view-search');
        //lVs=  ;
        this.searchNode = lSearchView;
        this.modes = this.action.view_mode.split(",")
        // 当有可用的View时才写换mode
        if (this.GetViewMode("view") && this.action.views && this.action.views.length > 0) {
            //this.mode = undefined;

            //如果存在且支持该Mode
            //  lMode = this.query["view"];// url View
            if (this.GetViewMode("view") && this.GetViewMode("view").indexOf(this.action.view_mode)) {
                this.SetViewMode(this.GetViewMode("view"));
            } else if (this.action.view_type) {
                //初始化Mode
                this.SetViewMode(this.action.view_type); // 更新Mode
            } else {
                //  this.viewId = action.views[0][1];
                this.SetViewMode(this.action.views[0][0]);
            }
        } else {
            //   this.SetViewMode( undefined);
        }

        //this.fire("on-cp-view-changed", this);
        this.dispatchEvent(new CustomEvent('on-cp-view-changed', this));
    }

    onActionChanged(action) {
        var b
        // b=a;
    }

    //-----------------------------------------
    // 1.根据Action更新而更新或隐藏
    onViewsChanged(views) {

    }

    onModeBtnClick(e) {
        // 改变Mode并激活事件
        this.SetViewMode(event.currentTarget.modeType);
    }

    onHotkeys(e) {
        // ESC
        if (e.keyCode === 27 && Polymer.dom(e).rootTarget === this.$.query) {
            this.showingSearch = false;
        }
    }

    // 变更Mode
    SetViewMode(mode) {
        if (mode != "") {
            //query = this.GetBaseQueryString();
            var query = this.GetQuery();
            // 只更新不同的Mode
            if (query["view"] != mode) {
                query["view"] = mode;
                //this.query = newMap;
                this.SetQuery(query);
            }

            // this.fire("on-cp-mode-changed", mode)
            this.dispatchEvent(new CustomEvent("on-cp-mode-changed", { mode: mode }));
        }
    }
}

customElements.define('app-view-control-panel', AppViewControlPanel);