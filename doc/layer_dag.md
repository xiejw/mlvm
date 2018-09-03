# Layer Graph Design

User builds a layer graph by calling APIs. All inputs are `InputLayer` while all
outputs are passed to `Graph` via explicit

## Compilation Phases

- Emit Layer DAG
- Lower to Op DAG
- Insert Send/Recv

