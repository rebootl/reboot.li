<script lang="ts">
  //import AnalogClock from '$lib/AnalogClock.svelte';
</script>

<div class="wrapper container">
  <div class="login-box">
    <h1>Chronolog</h1>
    <!--<div class="clock-box">
      <AnalogClock />
    </div>-->
    <p><em>Time tracking made simple.</em></p>
    <form method="POST" class="login-form" action="/login">
      <input type="text" name="username" placeholder="Username" />
      <input type="password" name="password" placeholder="Password" />
      <button>LOGIN</button>
    </form>
  </div>
</div>

<style>
/*  .wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .login-box {
    display: flex;
    flex-direction: column;
    align-items: center;
    max-width: 320px;
  }
  h1 {
    font-weight: 400;
    font-size: 2rem;
    letter-spacing: 0.8rem;
    padding-left: 0.6rem;
    text-align: center;
  }
  .clock-box {
    width: 80vw;
    height: 80vw;
  }
  .login-form {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 50vw;
    gap: 1rem;
  }*/
</style>
