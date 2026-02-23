# Exercise 4: AI as Learning Amplifier

## Topic: Wireless Routing Protocols

---

# Phase 1 — Build Foundation (My Own Understanding First)

## What I Understand About Wireless Routing Protocols

Wireless routing protocols determine how devices send data to each other without cables.
Because devices can move or lose connection, the routing system must dynamically adapt.

### Three Main Types of Wireless Routing Protocols

---

### 1. Proactive Protocols

These continuously maintain updated routing tables.

**Example:** Optimized Link State Routing (OLSR)

- All routes are known in advance
- Fast packet forwarding
- High overhead due to constant route updates

---

### 2. Reactive Protocols

Routes are created only when needed.

**Example:** Ad hoc On-Demand Distance Vector (AODV)

- Lower overhead
- Slight delay during route discovery
- Better suited for dynamic networks

---

### 3. Hybrid Protocols

Combine proactive and reactive approaches.

**Example:** Zone Routing Protocol (ZRP)

- Proactive within local zones
- Reactive for distant nodes

---

## Simple Scenario (8 Devices)

### Emergency Rescue Team Network

**Devices:**
- 6 rescue workers (mobile)
- 1 command vehicle (semi-stationary)
- 1 drone (high mobility)

**Total:** 8 devices

### My Protocol Choice: AODV

**Reasoning:**

- Nodes are mobile
- Network topology changes frequently
- Maintaining full routing tables (like OLSR) would waste bandwidth
- On-demand routing reduces overhead

**Conclusion:** Reactive routing fits better.

---

## Brief Reflection (Phase 1)

I explained everything without checking notes.
If I can explain it simply, I understand it.

---

# Phase 2 — Strategic AI Use

After forming my foundation, I asked targeted questions.

---

## Question 1
**What happens if node mobility increases dramatically in AODV?**

### AI Summary Response:
- Frequent route breakages
- High route discovery overhead
- Increased latency

### What I Learned

Reactive protocols are efficient,
but under extreme mobility, constant rediscovery becomes expensive.

There is a trade-off between:

- Control overhead
- Route stability
- Latency

---

## Question 2
**When would OLSR outperform AODV?**

### AI Explanation:
- In dense networks
- When traffic is constant
- When low latency is critical

### Insight

No routing protocol is universally best.

---

## Brief Reflection (Phase 2)

AI did not replace my understanding.
It exposed edge cases I had not considered.

It helped stress-test my reasoning.

---

# Phase 3 — Real Application

## Smart City Network Design

### Devices:
- 1,000 IoT sensors (mostly static)
- 50 traffic lights (static, critical)
- 10 emergency vehicles (mobile)

**Total:** 1,060 devices

---

## Initial Protocol Decisions

### IoT Sensors

- Mostly static
- Large number of nodes 

**Choice:** Proactive routing (e.g., OLSR)

**Reasoning:**
- Stable topology
- Low delay required
- Routes should already exist

---

### Traffic Lights

- Critical infrastructure
- High reliability required

**Choice:** Proactive routing and centralized monitoring

**Reasoning:**
- Low latency
- Predictable routing

---

### Emergency Vehicles

- High mobility

**Choice:** Reactive routing (e.g., AODV)

**Reasoning:**
- Vehicles move fast
- On-demand routing adapts better

---

## Potential Failure Points

- Congestion in dense sensor areas
- Route flooding during emergencies
- Mobility causing packet loss
- Single point of failure in central coordination
- Security risks (spoofed routing messages)

---

## Refinement with AI

### Question:
Would a hybrid protocol be better for 1,000 static sensors and 10 mobile vehicles?

### AI Suggestion:
- Hybrid zoning may reduce control overhead
- Local proactive clusters
- Reactive inter-zone routing

### Refined Design

Use a ZRP like hybrid structure:

- Sensors grouped into zones
- Proactive routing inside zones
- Reactive routing between zones and mobile units

This reduces global overhead while maintaining flexibility.

---

## Brief Reflection (Phase 3)

AI improved architectural efficiency.
But I structured the design first.

AI optimized — it did not originate.

---

# Final Reflection

## Human Judgment vs AI Contribution

- 70% my reasoning
- 30% AI refinement

AI helped explore edge cases and scaling implications.

---

## Could I Defend My Decisions Without AI?

Yes.

My reasoning is based on:

- Mobility
- Scalability
- Latency trade offs

AI only added optimization insight.

---

## What Will I Remember in 6 Months?

- Difference between proactive, reactive, and hybrid routing
- Trade-off between overhead and latency
- Mobility changes everything
- No protocol is universally best

---

## Did AI Make Me Sharper or Think for Me?

It made me sharper.

Because:

- I attempted first
- Then validated
- Then refined

If I had asked AI first, I would have copied instead of understood.

---

# Ethical / Learning Insight

AI should:

- Challenge assumptions
- Expose blind spots
- Simulate edge cases

AI should not:

- Replace reasoning
- Shortcut thinking

Learning amplification works only when human thinking comes first.