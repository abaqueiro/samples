# LVM - Logical Volume Manager

### List physical devices

```pvdisplay

  --- Physical volume ---
  PV Name               /dev/sda3
  VG Name               ubuntu-vg
  PV Size               <198.00 GiB / not usable 16.50 KiB
  Allocatable           yes
  PE Size               4.00 MiB
  Total PE              50687
  Free PE               22784
  Allocated PE          27903
  PV UUID               R6yYQD-5SLO-RyKZ-NZ3i-3vfY-VNQz-fDfCNR

```

### List volume groups

```vgdisplay

  --- Volume group ---
  VG Name               ubuntu-vg
  System ID
  Format                lvm2
  Metadata Areas        1
  Metadata Sequence No  6
  VG Access             read/write
  VG Status             resizable
  MAX LV                0
  Cur LV                2
  Open LV               2
  Max PV                0
  Cur PV                1
  Act PV                1
  VG Size               <198.00 GiB
  PE Size               4.00 MiB
  Total PE              50687
  Alloc PE / Size       27903 / <109.00 GiB
  Free  PE / Size       22784 / 89.00 GiB
  VG UUID               7UQlWm-2hX3-sdNK-sOwC-XGQo-R2gc-fpuIiK
```

### How to extend a L

