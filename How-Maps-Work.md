
- >= 1.24 version onwards, map in go uses swisstable algo

- Each key in a map forms 64 bit hash -> 64bit fits into 8 bytes
- each key is split into H1 and H2 --> 57bit and 7bits respectively

- The 64 bit hash is split into two parts H1 and H2
- Each Group contains 8 slots

- SwissTable 
    - Table
        - Group-1
            - ControlByte(8 slots of 8 byte)--> H2 of a key
            - Keys (8)
            - Values (8)
        - Group-2

- ControlByte -> 8 Bytes 
- Each byte of for H2 of a key
- So it can store 8 slots 
- Slot State ControlByte: empty, deleted or used/H2




