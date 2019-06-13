<template>
  <div id="app">
    <!-- default Vue CLI template -->
    <img alt="Vue logo"
         src="./assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App" />

    <hr>

    <h3>Problem Description</h3>
    <div class="description">
      <p class="text">This simple Vue.js app is broken. The app should allow the user to maintain a Justice League roster which displays a check mark next to the heroes chosen for the mission. The user should be able to build a team of at most three
        <span style="text-decoration: underline">unique</span> heroes for this mission. Adding a hero is accomplished by selecting an available hero using the drop down menu and then clicking the 'Add Hero' button. The user should be able to remove a hero from the chosen heroes list by clicking the 'x' in the top right corner of the box. Clicking the edit link next to a hero's name allows the user to change the name of that hero by typing a new name into the input box appearing below the roster and pressing enter. The input box disappears without making any changes if the element loses focus. The name change should propagate to any other mention of the hero in the app. Finally, there should be a 'Launch Mission' button next to the 'Add Hero' button that opens an alert validating the user's choice of exactly three heroes.</p>

      <!-- tasks -->
      <ul>
        <h3>Tasks:</h3>
        <li>Make the check mark appear
          <i>
            <strong>only</strong>
          </i> if this hero appears in the chosen heroes list.</li>
        <li>Prevent the user from adding more than three heroes to the chosen heroes list.</li>
        <li>Devise a way to alter the list of available heroes in the drop down whenever a hero is chosen to prevent the ability to add the same hero more than once.</li>
        <li>Ensure that when a hero is removed from the chosen heroes list, that hero is returned to the options in the drop down.</li>
        <li>Manage app state so that editing a hero's name propagates to all parts of the app. In case this isn't clear, edit Superman's name and observe that the name changes appropriately in the drop down. Refresh the browser to remove that change. Now add Superman to your list of chosen heroes and then edit his name and observe that the name does not update in the list of chosen heroes. This is not expected behavior and needs to be fixed.</li>
        <li>Add a 'Launch Mission' button next to the 'Add Hero' button that when clicked opens an alert stating either 'Mission complete' if the user has added exactly three heroes or 'We need three heroes' otherwise.</li>
        <li>While this app is simple enough to exist within a single Vue file this isn't an exercise in refactoring. Please accomplish the tasks without removing MissionPlanner.vue, ChosenHeroes.vue, or Hero.vue even if it would make state management easier. Feel free to add source files or move functionality/state from a parent component to its child component(s) or vice versa.</li>
      </ul>

      <!-- tips -->
      <ul>
        <h3>Tips:</h3>
        <li>You could implement this app using vanilla JS. You could also use jQuery or some other DOM manipulation library if you could figure out how to import it into an existing Vue application. However, this is not about the destination, it's about the journey. We are interested in your ability to learn
          <i>
            <strong>new</strong>
          </i> things. We prefer that you implement a 50% solution using Vue over a 100% solution using methods like
          <code>document.getElementById()</code> or
          <code>$()</code>.</li>
        <li>Vue ships with some helpful error handling, so as you develop, keep the console open to help with debugging.</li>
        <li>You may wish to use the Vue DevTools extension for Chrome.</li>
        <li>Any IDE or text editor can be used to make your changes, but VS Code and the Vetur extension for VS Code go a long way towards making Vue components easy to write and debug.</li>
        <li>Some of the app's functionality is already in place, we suggest clicking around to figure out what works and what does not and why as a first step.</li>
        <li>Think about solving the problem of state management, then write the code. The check marks need to <i>know about</i> the list of chosen heroes. The list of chosen heroes needs to <i>know about</i> changes to a hero's name. Getting all three components on the same page is the real challenge you need to solve.</li>
        <li>None of the application state is persistent, so refreshing the browser will get you back to square one.</li>
        <li>The existing code is only a starting point. It is intentionally poorly written and designed. Replace it as you see fit with your own code to accomplish the objectives. For example, don't get locked into thinking a hero is an object that only has a 'name' property just because the current code employs that model. You should add whatever properties you think you need to get the job done.</li>
        <li>We expect that you've never worked with Vue before. Don't be intimidated by it. The documentation (see links above) is great and nothing you are being asked to do is fundamentally different from the object-oriented programming, loops, and conditional statements you are familiar with.</li>
        <li>The only three files you might need to edit are MissionPlanner.vue, ChosenHeroes.vue, and Hero.vue. Everything else in this directory was generated by the Vue CLI (a scaffolding tool for Vue projects) and can safely be ignored.</li>
        <li>Use JavaScript's alert() method to implement the click handler for the 'Launch Mission' button. No need to get fancy.</li>
        <li>If any of the objectives are unclear, let us know as soon as possible so that we can clarify. Good luck!</li>
      </ul>
    </div>

    <hr>

    <!-- initial mission planner app -->
    <MissionPlanner/>

  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import MissionPlanner from "./MissionPlanner.vue";

export default {
  name: "app",
  components: {
    HelloWorld,
    MissionPlanner
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  margin-bottom: 100px;
}
.description {
  text-align: left;
}
.description .text {
  margin-left: 2em;
  margin-right: 2em;
}
</style>
