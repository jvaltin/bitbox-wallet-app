/**
 * Copyright 2018 Shift Devices AG
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import preactCliTypeScript from 'preact-cli-plugin-typescript';

/**
 * Function that mutates original webpack config.
 * Supports asynchronous changes when promise is returned.
 *
 * @param {object} config - original webpack config.
 * @param {object} env - options passed to CLI.
 * @param {WebpackConfigHelpers} helpers - object with useful helpers when working with config.
 */
export default function (config, env, helpers) {
    if (!env.production) {
        config.devServer.overlay = true;
        config.devServer.hot = false;
    }

    if (env.production) {
        helpers.getPluginsByName(config, 'UglifyJsPlugin')[0].plugin.options.sourceMap = false;
    }

    // Support for TypeScript (https://github.com/wub/preact-cli-plugin-typescript):
    preactCliTypeScript(config);

    // In order to import CSS Modules (https://github.com/css-modules/css-modules) from TypeScript,
    // we need to replace the existing CSS Loader (https://github.com/webpack-contrib/css-loader)
    // with https://github.com/Jimdo/typings-for-css-modules-loader, which generates '*.css.d.ts'
    // files that export the class names as strings. With TypeScript, you can no longer reference
    // CSS classes that do not exist. Use `import * as style from './styles.css';` to import them.

    // Retrieve the rules in which the CSS Loader is used. The CSS Loader is used twice: The first
    // time only for CSS (and other) files in 'src/components' and 'src/routes', the second time for
    // CSS (and other) files everywhere else. This distinction is important because the CSS classes
    // are only renamed to avoid name collisions in the former but not the latter rule. This is why
    // the styles imported in 'src/index.js' can be used globally. Alternatively, the loaders could
    // be suppressed there with `import '!style-loader!css-loader!./style';`. (If you want a single
    // CSS rule to be available globally, you can set its scope with `:global(.className) {…}` (see
    // https://github.com/css-modules/css-modules#exceptions).) We only want to replace the first:
    const rule = helpers.getLoadersByName(config, 'css-loader')[0];

    config.module.loaders[rule.ruleIndex].loader[rule.loaderIndex] = {
        loader: 'typings-for-css-modules-loader',
        options: {
            // Instead of exporting CSS class names as properties of an interface (allowing dashes),
            // export the names directly and transform dashes to valid variable names with camelCase
            // (see https://github.com/Jimdo/typings-for-css-modules-loader#namedexport-option):
            camelCase: true,
            namedExport: true,
            // Existing options (see https://github.com/Jimdo/typings-for-css-modules-loader#usage):
            modules: true,
            sourceMap: false,
            importLoaders: 1,
            // See https://github.com/webpack/loader-utils#interpolatename for more options:
            localIdentName: '[path][name]-[local]', // Was '[local]__[hash:base64:5]' before.
        }
    };
}
