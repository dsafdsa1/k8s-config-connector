{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ComputeServiceAttachment{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Compute Engine</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/compute/docs/">/compute/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>beta.serviceAttachments</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/compute/docs/reference/rest/beta/serviceAttachments">/compute/docs/reference/rest/beta/serviceAttachments</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpcomputeserviceattachment<br>gcpcomputeserviceattachments<br>computeserviceattachment</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>compute.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>computeserviceattachments.compute.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties



### Spec
#### Schema
  ```yaml
  connectionPreference: string
  consumerAcceptLists:
  - connectionLimit: integer
    projectRef:
      external: string
      name: string
      namespace: string
  consumerRejectLists:
  - external: string
    name: string
    namespace: string
  description: string
  enableProxyProtocol: boolean
  location: string
  natSubnets:
  - external: string
    name: string
    namespace: string
  projectRef:
    external: string
    name: string
    namespace: string
  resourceID: string
  targetServiceRef:
    external: string
    name: string
    namespace: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>connectionPreference</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The connection preference of service attachment. The value can be set to `ACCEPT_AUTOMATIC`. An `ACCEPT_AUTOMATIC` service attachment is one that always accepts the connection from consumer forwarding rules. Possible values: CONNECTION_PREFERENCE_UNSPECIFIED, ACCEPT_AUTOMATIC, ACCEPT_MANUAL{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Projects that are allowed to connect to this service attachment.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists[].connectionLimit</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The value of the limit to set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists[].projectRef</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists[].projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The project id or number for the project to set the limit for.

Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists[].projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerAcceptLists[].projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerRejectLists</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerRejectLists[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerRejectLists[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerRejectLists[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>consumerRejectLists[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}An optional description of this resource. Provide this property when you create the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>enableProxyProtocol</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Immutable. If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The location for the resource{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>natSubnets</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>natSubnets[]</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>natSubnets[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeSubnetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>natSubnets[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>natSubnets[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. The Project that this resource belongs to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The project for the resource

Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The URL of a service serving the endpoint identified by this service attachment.

Allowed value: The `selfLink` field of a `ComputeForwardingRule` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>targetServiceRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>{% verbatim %}* Field is required when parent field is specified{% endverbatim %}</p>


### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  connectedEndpoints:
  - endpoint: string
    pscConnectionId: integer
    status: string
  fingerprint: string
  id: integer
  observedGeneration: integer
  pscServiceAttachmentId:
    high: integer
    low: integer
  region: string
  selfLink: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>connectedEndpoints</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}An array of connections for all the consumers connected to this service attachment.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>connectedEndpoints[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>connectedEndpoints[].endpoint</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The url of a connected endpoint.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>connectedEndpoints[].pscConnectionId</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The PSC connection id of the connected endpoint.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>connectedEndpoints[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The status of a connected endpoint to this service attachment. Possible values: PENDING, RUNNING, DONE{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>fingerprint</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Fingerprint of this resource. This field is used internally during updates of this resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>id</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The unique identifier for the resource type. The server generates this identifier.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>pscServiceAttachmentId</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}An 128-bit global unique ID of the PSC service attachment.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>pscServiceAttachmentId.high</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>pscServiceAttachmentId.low</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>region</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>selfLink</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Server-defined URL for the resource.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
  ```yaml
  # Copyright 2021 Google LLC
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
  
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeServiceAttachment
  metadata:
    name: computeserviceattachment-sample
  spec:
    projectRef:
       # Replace ${PROJECT_ID?} with your project ID
       external: "projects/${PROJECT_ID?}"
    location: "us-west2"
    description: "A sample service attachment"
    targetServiceRef:
      name: "computeserviceattachment-dep"
    connectionPreference: "ACCEPT_MANUAL"
    natSubnets:
    - name: "computeserviceattachment-dep1"
    enableProxyProtocol: false
    consumerRejectLists:
    - name: "computeserviceattachment-dep1"
    consumerAcceptLists:
    - projectRef:
        name: "computeserviceattachment-dep2"
      connectionLimit: 2
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeBackendService
  metadata:
    name: computeserviceattachment-dep
  spec:
    location: "us-west2"
    networkRef:
      name: "computeserviceattachment-dep"
    loadBalancingScheme: "INTERNAL"
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeForwardingRule
  metadata:
    name: computeserviceattachment-dep
  spec:
    location: "us-west2"
    networkRef:
      name: "computeserviceattachment-dep"
    subnetworkRef:
      name: "computeserviceattachment-dep3"
    description: "A test forwarding rule with internal load balancing scheme"
    loadBalancingScheme: "INTERNAL"
    backendServiceRef:
      name: "computeserviceattachment-dep"
    networkTier: "PREMIUM"
    allPorts: true
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeNetwork
  metadata:
    name: computeserviceattachment-dep
  spec:
    autoCreateSubnetworks: false
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeSubnetwork
  metadata:
    name: computeserviceattachment-dep1
  spec:
    region: "us-west2"
    ipCidrRange: "10.2.0.0/16"
    networkRef:
      name: "computeserviceattachment-dep"
    purpose: "PRIVATE_SERVICE_CONNECT"
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeSubnetwork
  metadata:
    name: computeserviceattachment-dep2
  spec:
    region: "us-west2"
    ipCidrRange: "10.3.0.0/16"
    networkRef:
      name: "computeserviceattachment-dep"
    purpose: "PRIVATE_SERVICE_CONNECT"
  ---
  apiVersion: compute.cnrm.cloud.google.com/v1beta1
  kind: ComputeSubnetwork
  metadata:
    name: computeserviceattachment-dep3
  spec:
    region: "us-west2"
    ipCidrRange: "10.4.0.0/16"
    networkRef:
      name: "computeserviceattachment-dep"
    purpose: "PRIVATE"
  ---
  apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
  kind: Project
  metadata:
    name: computeserviceattachment-dep1
  spec:
    organizationRef:
      # Replace "${ORG_ID?}" with the numeric ID for your organization
      external: "${ORG_ID?}"
    name: "computeserviceattachment-dep1"
  ---
  apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
  kind: Project
  metadata:
    name: computeserviceattachment-dep2
  spec:
    organizationRef:
      # Replace "${ORG_ID?}" with the numeric ID for your organization
      external: "${ORG_ID?}"
    name: "computeserviceattachment-dep2"
  ```


{% endblock %}