dimm-edac-mon
---

## Monitor Dimms for Error Detection And Correction events on Linux 

### Usage:
```sh
make
./dimm-edac-mon
```

### Example:
```sh
$ ./dimm-edac-mon
/sys/devices/system/edac/mc/mc0
ce_count: 0
ce_noinfo_count: 0
mc_name: Broadwell SrcID#0_Ha#0
seconds_since_reset: 3195552
size_mb: 32768
ue_count: 0
ue_noinfo_count: 0

/sys/devices/system/edac/mc/mc0/dimm0
dimm_ce_count: 0
dimm_dev_type: x4
dimm_edac_mode: S4ECD4ED
dimm_location: channel 0 slot 0
dimm_mem_type: Registered-DDR4
dimm_ue_count: 0
size: 8192

/sys/devices/system/edac/mc/mc0/dimm3
dimm_ce_count: 0
dimm_dev_type: x4
dimm_edac_mode: S4ECD4ED
dimm_location: channel 1 slot 0
dimm_mem_type: Registered-DDR4
dimm_ue_count: 0
size: 8192

/sys/devices/system/edac/mc/mc0/dimm6
dimm_ce_count: 0
dimm_dev_type: x4
dimm_edac_mode: S4ECD4ED
dimm_location: channel 2 slot 0
dimm_mem_type: Registered-DDR4
dimm_ue_count: 0
size: 8192

/sys/devices/system/edac/mc/mc0/dimm9
dimm_ce_count: 0
dimm_dev_type: x4
dimm_edac_mode: S4ECD4ED
dimm_location: channel 3 slot 0
dimm_mem_type: Registered-DDR4
dimm_ue_count: 0
size: 8192
```

### To-do:
- add support for timeseries & json output
- add nagios check
