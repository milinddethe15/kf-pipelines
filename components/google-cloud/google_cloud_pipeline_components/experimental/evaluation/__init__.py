# Copyright 2021 The Kubeflow Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Model evaluation components."""

from google_cloud_pipeline_components.experimental.evaluation.data_bias.component import detect_data_bias as DetectDataBiasOp
from google_cloud_pipeline_components.experimental.evaluation.model_bias.component import detect_model_bias as DetectModelBiasOp

__all__ = [
    'DetectModelBiasOp',
    'DetectDataBiasOp',
]
