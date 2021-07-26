<template>
  <div class="viewport">
    <div class="viewport-middle-section">
      <div class="ui form jug-riddle-form">
        <div class="fields">
          <div class="field">
            <label>X Jug</label>
            <input
              v-model="xCapacity"
              @input="clearError"
              type="number"
              min="1"
              placeholder="Capacity"
            />
          </div>
          <div class="field">
            <label>Y Jug</label>
            <input
              v-model="yCapacity"
              @input="clearError"
              type="number"
              min="1"
              placeholder="Capacity"
            />
          </div>
          <div class="field">
            <label>Goal</label>
            <input
              v-model="zGoal"
              @input="clearError"
              type="number"
              placeholder="Capacity"
            />
          </div>
        </div>
        <div class="field">
          <span class="error-form" v-if="showError"
            >Error: {{ errorMessage }}
          </span>
        </div>
        <div class="ui button" @click="obtainSolution">Obtain solution</div>
      </div>

      <div class="play-division" v-if="thereIsSolution">
        <button class="ui button" @click="startSimulation">
          <i class="play huge circle icon play-icon"></i>
        </button>
        <button class="ui button" @click="stopSimulation">
          <i class="stop huge circle icon play-icon"></i>
        </button>
      </div>
      
      <div class="no-solution-division" v-if="solution == null">
          <span> There is no solution for this inputs </span>
      </div>

      <div class="jugs-images">
        <div class="jug-image">
          <img class="ui medium image" :src="xJugImage" />
          <h2>Current Volume: {{ xCurrentVolume }}</h2>
        </div>
        <div class="jug-image">
          <img class="ui medium image" :src="yJugImage" />
          <h2>Current Volume: {{ yCurrentVolume }}</h2>
        </div>
      </div>
      <div class="leak-image">
        <img class="ui image" :src="lake" />
      </div>

      <ol class="solution-steps">
        <li
          class="step"
          @click="moveToStep(index)"
          :class="{ highlight: currentStep === index }"
          v-for="(step, index) in solution"
          v-bind:key="index"
        >
          <span>X:</span> {{ step.x }}. <span>Y:</span> {{ step.y }}
        </li>
      </ol>
    </div>
  </div>
</template>

<script>
import empty from "../../assets/jug-empty.png";
import half from "../../assets/jug-half.png";
import full from "../../assets/jug-full.png";
import lake from "../../assets/lake.png";
import riddelsController from "../controllers/riddles";
export default {
  name: "Home",
  components: {},
  mounted() {},
  data() {
    return {
      visible: false,
      lake,
      xCurrentVolume: 0,
      yCurrentVolume: 0,
      xCapacity: null,
      yCapacity: null,
      zGoal: null,
      solution: [],
      showError: false,
      errorMessage: "",
      currentStep: null,
      interval: null,
    };
  },
  computed: {
    thereIsSolution() {
      if (!this.solution)  {
        return false;
      }
      return this.solution.length > 0;
    },
    xJugImage() {
      if (this.xCurrentVolume == 0) return empty;
      else if (this.xCurrentVolume == this.xCapacity) return full;
      return half;
    },
    yJugImage() {
      if (this.yCurrentVolume == 0) return empty;
      else if (this.yCurrentVolume == this.yCapacity) return full;
      return half;
    },
  },
  methods: {
    moveToStep(index) {
      const step = this.solution[index];
      this.xCurrentVolume = step.x;
      this.yCurrentVolume = step.y;
      this.currentStep = index;
    },
    stopSimulation() {
      clearInterval(this.interval);
      this.interval = null;
    },
    startSimulation() {
      if (this.interval !== null) {
        return;
      }
      this.currentStep = 0;
      this.xCurrentVolume = this.solution[this.currentStep].x;
      this.yCurrentVolume = this.solution[this.currentStep].y;
      this.interval = setInterval(() => {
        this.currentStep += 1;
        if (this.currentStep >= this.solution.length) {
          clearInterval(this.interval);
          this.currentStep = null;
          this.interval = null;
          return;
        }

        this.xCurrentVolume = this.solution[this.currentStep].x;
        this.yCurrentVolume = this.solution[this.currentStep].y;
      }, 1000);
    },
    validParameters() {
      if (this.xCapacity <= 0) {
        this.errorMessage = "X jug capacity must be greater than 0";
        this.showError = true;
        return false;
      }
      if (this.yCapacity <= 0) {
        this.errorMessage = "Y jug capacity must be greater than 0";
        this.showError = true;
        return false;
      }
      if (this.zGoal < 0) {
        this.errorMessage = "The goal cannot be negative";
        this.showError = true;
        return false;
      }
      return true;
    },
    clearError() {
      this.showError = false;
      this.errorMessage = "";
    },
    async obtainSolution() {
      this.clearError();
      if (!this.validParameters()) {
        return;
      }
      try {
        this.solution = await riddelsController.obtainJugRiddleSolution(
          this.xCapacity,
          this.yCapacity,
          this.zGoal
        );
      } catch (err) {
        console.log(`err: `, err);
        this.showError = true;
        this.errorMessage = err;
      }
    },
  },
};
</script>
<style>
.viewport {
  background-repeat: no-repeat;
  background-color: rgb(61, 61, 61);
  color: white;
  display: flex;
  height: 100%;
  width: 100%;
  overflow: auto;
}
.viewport-middle-section {
  display: flex;
  background-color: transparent;
  width: 700px;
  margin-left: auto;
  margin-right: auto;
  border: black 1px solid;
  background-color: rgba(158, 163, 160, 255);
  flex-direction: column;
  overflow: auto;
}
.jug-riddle-form {
  margin-left: auto;
  margin-right: auto;
  margin-top: 50px;
}

.jugs-images {
  display: flex;
}
.jug-image {
  margin: auto;
  margin-top: 40px;
  display: flex;
  flex-direction: column;
}
.jug-image h2 {
  margin: auto;
}
.leak-image {
  display: flex;
}
.leak-image img {
  width: 90%;
  margin: auto;
  margin-top: 40px;
}
.error-form {
  background-color: transparent;
  color: red;
  font-weight: bold;
}

.solution-steps {
  width: 90%;
  min-height: 150px;
  margin: auto;
  margin-top: 40px;
  margin-bottom: 100px;
  overflow-y: auto;
}

.step {
  padding: 3px;
  margin: 5px;
}
.step span {
  font-weight: bold;
}
.highlight {
  border: 2px solid red;
}

.play-division {
  width: 90%;
  margin: auto;
  margin-top: 40px;
  display: flex;
}
.play-division .ui.button {
  margin: auto;
  background-color: transparent !important;
}

.no-solution-division {
  width: 90%;
  margin: auto;
  margin-top: 40px;
  display: flex;
}

.no-solution-division span {
  margin: auto;
  font-weight: bold;
  font-size: 16px;
  color: orange;
}

.play-icon {
  margin: auto !important;
}
</style>
