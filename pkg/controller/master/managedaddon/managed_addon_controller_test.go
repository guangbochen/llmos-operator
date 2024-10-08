package managedaddon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeValuesContent(t *testing.T) {
	type input struct {
		key          string
		values       string
		customValues string
	}
	type output struct {
		values string
		err    error
	}

	var testCases = []struct {
		name     string
		given    input
		expected output
	}{
		{
			name: "merge values",
			given: input{
				key: "test",
				values: `
operatorNamespace: storage-system

clusterName: llmos-ceph

configOverride:
# configOverride: |
#   [global]
#   mon_allow_pool_delete = true
#   osd_pool_default_size = 3
#   osd_pool_default_min_size = 2

cephClusterSpec:
  cephVersion:
    image: quay.io/ceph/ceph:v18.2.4

  dataDirHostPath: /var/lib/llmos/rook

  mon:
    count: 1 # default to 1 on cluster-init
    allowMultiplePerNode: false

  mgr:
    count: 1 # default to 1 on cluster-init
    allowMultiplePerNode: false

  placement:
    all:
      nodeAffinity:
        preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 1
          preference:
            matchExpressions:
            - key: role
              operator: In
              values:
              - storage
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/control-plane
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/storage
        operator: Exists

  resources:
    mgr:
      limits:
        memory: "1Gi"
      requests:
        cpu: "500m"
        memory: "512Mi"
    mon:
      limits:
        memory: "2Gi"
      requests:
        cpu: "1000m"
        memory: "1Gi"
    osd:
      limits:
        memory: "4Gi"
      requests:
        cpu: "1000m"
        memory: "2Gi"
    prepareosd:
      requests:
        cpu: "500m"
        memory: "50Mi"
    mgr-sidecar:
      limits:
        memory: "100Mi"
      requests:
        cpu: "100m"
        memory: "40Mi"
    crashcollector:
      limits:
        memory: "60Mi"
      requests:
        cpu: "100m"
        memory: "60Mi"
    logcollector:
      limits:
        memory: "1Gi"
      requests:
        cpu: "100m"
        memory: "100Mi"
    cleanup:
      limits:
        memory: "1Gi"
      requests:
        cpu: "500m"
        memory: "100Mi"
    exporter:
      limits:
        memory: "128Mi"
      requests:
        cpu: "50m"
        memory: "50Mi"

# -- A list of CephBlockPool configurations to deploy
# @default -- See [below](#ceph-block-pools)
cephBlockPools:
  - name: llmos-ceph-blockpool
    spec:
      failureDomain: host
      replicated:
        size: 2 # default to 2 on cluster init
      parameters:
        min_size: "1" # default to 1 on cluster init
`,
				customValues: `
operatorNamespace: storage-system
clusterName: llmos-ceph
configOverride: |
  [global]
  mon_allow_pool_delete = true
  osd_pool_default_size = 3
  osd_pool_default_min_size = 2

cephClusterSpec:
  cephVersion:
    image: quay.io/ceph/ceph:v19.2.4
  dataDirHostPath: /var/lib/llmos/rook

  mon:
    count: 2 # default to 1 on cluster-init
    allowMultiplePerNode: false

  mgr:
    count: 2 # default to 1 on cluster-init
    allowMultiplePerNode: false

  resources:
    mgr:
      limits:
        memory: "1Gi"
      requests:
        cpu: "1000m"
        memory: "1Gi"
    mon:
      limits:
        memory: "2Gi"
      requests:
        cpu: "1000m"
        memory: "1Gi"
    osd:
      limits:
        memory: "4Gi"
      requests:
        cpu: "2000m"
        memory: "4Gi"

# -- A list of CephBlockPool configurations to deploy
# @default -- See [below](#ceph-block-pools)
cephBlockPools:
  - name: llmos-ceph-blockpool
    spec:
      failureDomain: host
      replicated:
        size: 3 # default to 2 on cluster init
      parameters:
        min_size: "2" # default to 1 on cluster init
`,
			},
			expected: output{
				values: `cephBlockPools:
    - name: llmos-ceph-blockpool
      spec:
        failureDomain: host
        parameters:
            min_size: "2"
        replicated:
            size: 3
cephClusterSpec:
    cephVersion:
        image: quay.io/ceph/ceph:v19.2.4
    dataDirHostPath: /var/lib/llmos/rook
    mgr:
        allowMultiplePerNode: false
        count: 2
    mon:
        allowMultiplePerNode: false
        count: 2
    resources:
        mgr:
            limits:
                memory: 1Gi
            requests:
                cpu: 1000m
                memory: 1Gi
        mon:
            limits:
                memory: 2Gi
            requests:
                cpu: 1000m
                memory: 1Gi
        osd:
            limits:
                memory: 4Gi
            requests:
                cpu: 2000m
                memory: 4Gi
clusterName: llmos-ceph
configOverride: |
    [global]
    mon_allow_pool_delete = true
    osd_pool_default_size = 3
    osd_pool_default_min_size = 2
operatorNamespace: storage-system
`,
				err: nil,
			},
		},
	}

	for _, tc := range testCases {
		var actual output
		actual.values, actual.err = mergeValues(tc.given.values, tc.given.customValues)
		assert.Equal(t, tc.expected, actual)
	}
}
