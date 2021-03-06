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

import { h, RenderableProps } from 'preact';
import { translate } from 'react-i18next';
import { Translate } from '../../utils/translate';
import loading from '../../decorators/loading';
import Status from '../status/status';
import A from '../anchor/anchor';

/**
 * Describes the file that is loaded from 'https://shiftcrypto.ch/updates/desktop.json'.
 */
interface File {
    current: string;
    version: string;
    description: string;
}

interface Props {
    t: Translate;
    file: File | null;
}

function Update({t, file}: RenderableProps<Props>): JSX.Element | null {
    return file && (
        <Status dismissable keyName={`update-${file.version}`} type="info">
            {t('app.upgrade', {
                current: file.current,
                version: file.version
            })}
            {file.description}
            {' '}
            <A href="https://shiftcrypto.ch/start">
                {t('button.download')}
            </A>
        </Status>
    );
}

const LoadingUpdate = translate()(loading<Props>({ file: 'update' })(Update));
export default LoadingUpdate;
