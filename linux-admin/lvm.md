# LVM - Logical Volume Manager

### List physical devices

```
pvdisplay

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

```
vgdisplay

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

### List logical volumes

```
lvdisplay

  --- Logical volume ---
  LV Path                /dev/ubuntu-vg/ubuntu-lv
  LV Name                ubuntu-lv
  VG Name                ubuntu-vg
  LV UUID                dkZX3D-WdEx-2okU-uNa4-pnXI-JPkN-8iamto
  LV Write Access        read/write
  LV Creation host, time ubuntu-server, 2023-10-13 16:17:09 +0000
  LV Status              available
  # open                 1
  LV Size                <49.00 GiB
  Current LE             12543
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     256
  Block device           253:0

  --- Logical volume ---
  LV Path                /dev/ubuntu-vg/container-images
  LV Name                container-images
  VG Name                ubuntu-vg
  LV UUID                s4qbKr-h11X-RrBK-IJRv-MbZe-qRVy-02Tj73
  LV Write Access        read/write
  LV Creation host, time server02.local, 2024-03-26 16:08:06 +0000
  LV Status              available
  # open                 1
  LV Size                60.00 GiB
  Current LE             15360
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     256
  Block device           253:1

```

### How to extend a Logical volume

```
# extend the logical volume
lvextend -L +20G -r /dev/ubuntu-vg/container-images

# resize the filesystem in the volume (this case is an ext2 filesystem)
resize2fs /dev/mapper/ubuntu--vg-container--images
```

