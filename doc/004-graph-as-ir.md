Topic: Graph As IR
==================

- Draft: 2019-09
- State: Drafting

Zero: Why Graph As IR?
----------------------

Graph, as IR,

- captures the dependency by default,
- is easy for cluster analysis,
- is easy to be executed in parallel.
- is beautify for visualization.


First: Pure Data-Flow Graph
---------------------------

Pure data-flow graph is the graph without side-effects, like variable writes,
resources updates/reads.

It is easy to understand and straightforward to transform.

- It should be very easy and keep that way. It reflects the math we care.
- Transformation should be also easy. Example like symbolic gradient is much
  safe and unchallenging to be performed on pure data-flow graph.

Next: Adding Variable Write
---------------------------

Adding the support for variable write is not trivial. See design doc 001, 002,
and 003 in this folder for context and discussions.

002 was chosen due to simplicity and safeness. However, the first implementation
was still too complicated.

### The v1 Implementation

The v1 implementation is simple

1. Introduced a `init` Node for variable initialization.
2. All Nodes, using variables as operands, depend on the `init` Node.
3. We am aiming to make sure that variable updates are scheduled _after_ all
   usages.

In particular, toward bullet 3, some early design was implemented:

- Some control Nodes were introduced. Their responsibility were purely for
  dependency management in Graph. For example, all output Nodes have edges to a
  control Node, called `ctl`.
- All variable update Nodes have edges from `ctl`. The hope is the separating
  the computation with variable updates for safeness.

Much later, it was realized

1. **Graph rewrite is painful**. How/When to attach the "invisible" Nodes, like
   `init` and `ctl` above? What's the rule of thumb for this? How can we prove
   the rewrite could be correct?

2. **Not future proof**. It is quite foreseeable that adding any new stuff makes
   the combinations hard. Unless we can prove it is safe, it is always unsafe.
   For example, the `ctl` draws a line between outputs and variable updates. But
   what if the variable updates are using the results from a separated
   sub-graphs. This is not safe but we forget to validate this during design
   stage.

### Proposed Solution: Keep it Simple

Let's keep it simple, which will be future proof.

New design of the Graph structure

    type Graph struct {
        PreActions       map[*Action]bool // Set of actions. No order defined.
        Computation      *DataFlowGraph
        PostActions      map[*Action]bool // Set of actions. No order defined.
    }

In summary, we clearly define the boundary and maintain the simple data-flow
graph without control nodes.

