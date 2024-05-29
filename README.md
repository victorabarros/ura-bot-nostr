<p align="center">
  <img width="70%" src="https://raw.githubusercontent.com/victorabarros/ura-bot/main/assets/UraBot_profile01.png" />
</p>

# UraBot NOSTR

Page dedicated to sharing uranium market stock prices, analyses and relevant news.

## How to run

## Endpoints

## Flow

```mermaid
    flowchart LR;

      %% components
      cron[cron];
      server[UraBot<br>Server];
      finhub[finhub];
      client[nostr];

      %% flow
      cron--->
      server--->
      client;

      finhub----
      server;

      %% styles
      classDef box fill-opacity:.5, stroke:grey, stroke-width:.5px;
      class cron,server,finhub,db,client box
```

## Support

I'm more than happy to be honored with your support.

<p>
  <a href="https://www.buymeacoffee.com/victorbarros" target="_blank">
    <img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="45px">
  </a>
</p>

<p align="center">
  <br/>
  Made in Brazil
  <br/>
  <img src="https://user-images.githubusercontent.com/42843223/222024964-9494cd55-849c-40a3-8121-8fa00d575475.png" height="30px"/>
</p>
