<script lang="ts">
  import Spacer from './components/Spacer.svelte'
  import LayoutGrid, { Cell, InnerGrid } from '@smui/layout-grid'
  import Button from '@smui/button'
  import { peers, getPeers } from './stores/peerStore'
  import { offers, refreshOffers } from './stores/offerStore'
  import { connectAccount, currentAccount } from './stores/metamask'
  import OffersTable from './components/OffersTable.svelte'
  import StatCard from './components/StatCard.svelte'
  import TakeDealDrawer from './components/TakeDealDialog.svelte'

  const handleRefreshClick = () => {
    getPeers()
  }

  const connectMetamask = () => {
    connectAccount()
  }
</script>

<main>
  <LayoutGrid>
    <Spacer />
    <Cell spanDevices={{ desktop: 8, tablet: 6, phone: 12 }}>
        <Cell spanDevices={{ desktop: 8, tablet: 6, phone: 12 }}>
          <StatCard title="ETH-XMR Atomic Swap" content="Please ensure your Metamask is unlocked and set to the correct network before swapping. DO NOT REFRESH THE PAGE WHILE A SWAP IS HAPPENING!" />
        </Cell>
        <br />
      <InnerGrid>
        <Cell spanDevices={{ desktop: 2, tablet: 4, phone: 12 }}>
          <StatCard title="Peers" content={$peers.length.toString()} />
        </Cell>
        <Cell spanDevices={{ desktop: 2, tablet: 4, phone: 12 }}>
          <StatCard title="Offers" content={$offers.length.toString()} />
        </Cell>
        <Cell class="refreshButton">
          <Button on:click={handleRefreshClick}>Refresh</Button>
        </Cell>
        <Cell class="metamask">
          {#if $currentAccount}
            <StatCard title="Account" content={$currentAccount} />
          {:else}
            <Button on:click={connectMetamask}>Connect Metamask</Button>
          {/if}
        </Cell>
      </InnerGrid>
      <br />
      <OffersTable />
    </Cell>
    <TakeDealDrawer />
  </LayoutGrid>
</main>

<svelte:head>
  <!-- <link rel="stylesheet" href="node_modules/svelte-material-ui/bare.css" /> -->
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/svelte-material-ui@6.0.0-beta.13/bare.min.css"
  />
  <!-- Material Icons -->
  <link
    rel="stylesheet"
    href="https://fonts.googleapis.com/icon?family=Material+Icons"
  />
  <!-- Roboto -->
  <link
    rel="stylesheet"
    href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,600,700"
  />
  <!-- Roboto Mono -->
  <link
    rel="stylesheet"
    href="https://fonts.googleapis.com/css?family=Roboto+Mono"
  />
</svelte:head>

<style>
  * :global(.refreshButton) {
    display: flex;
    align-items: center;
  }
</style>
