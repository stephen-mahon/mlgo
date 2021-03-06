# Programming Exercise 1: Linear Regression
### Machine Learning

## 1 Linear regression with one variable
In this part of this exercise, you will implement linear regression with one variable to predict profits for a food truck.
Suppose you are the CEO of a restrauant franchise and are considering different cities for opening a new outlet.
The chain already has trucks in various cities and you have data for profits and populations from the cities.

You would like to use this data to help you select which city to expand to next.

The file `ex1data1.txt` contains the dataset for our linear regression problem. The first column is the population of a city and the second column is the profit of a food truck in that city. A negative value for profit indicates a loss.

The `ex1.go` script has already been set up to load this data for you.

### 1.1 Plotting the Data
Before starting on any task, it is often useful to understand the data by visualizing it.
For this dataset, you can use a scatter plot to visualize the data, since it has only two properties to plot (profit and population).
(Many other problems that you will encounter in real life are multi-dimensional and can't be plotted on a 2-d plot.)

In `ex1.m`, the dataset is loaded from the data file into the variable `xys`:

```
xys, err := readData(fileName) // filename = "ex1data1.txt"
if err != nil {
    log.Fatalf("could not read %v: %v", fileName, err)
}
```

Next, the script calls the fuction plotData to create a scatter plot of the data.
Your job is to complete `plotData.go` to draw the plot.

### 1.2 Gradient Descent
In this part, you will fit the linear regression paramaters 'theta' to our data set using gradient descent.

#### 1.2.1 Update Equations
The objective of linear regression is to minimize the cost function.
```
J(\theta) = 1/2m sum_{i=1}^{m}(h_\theta(x^{(i)})-y^{(i)})^2
```

where the hypothesis $$h_\theta(x)$$ is given by the linear model
```
h_\theta = \theta^T x = \theta_0 + \theta_1 x_1
```

Recall that the parameters of your model are the $$\theta_j$$ values.
These are the values you will adjust to minimize cost $$J(\theta)$$.
One way to do this is to use the batch gradient descent algorithm.
In batch gradient descent, each iteration performs the update


```
\theta_j := \theta_j - \alpha 1/m \sum_{i=1}^m(h_\theta(x^{(i)})-y^{(i)})x_j^{(i)} (simultaneously update all \theta_j for all j)
```

With each step of gradient descent, your parameters $$\theta_j$$ come closer to the
optimal values that will achieve the lowest cost $$J(\theta)$$.

#### 1.2.2 Implementation
in `ex1.go`, we have already set up the data for linear regression
In the following lines, we add another dimension to our data to accommodate the `\theta_0` intercept term.
We also initialize the initial parameters to 0 and the learning rate alpha is set as default to 0.01.

#### 1.2.3 Computing the Cost `J(theta)`
As you perform gradient descent to learn minimize the cost function `J(theta)`, it is helpful to monitor the convergence by computing the cost.
In this section, you will implement a function to calculate `J(theta)` so you can check the convergence of your gradient descent implementation.

Your next task is to complete the code in the file `ComputeCost.go`, which is a function that computes `J(theta)`.

Once you have completed the function, the next step in `ex1.go` will run `ComputeCost` once using `theta` initialized to zeros, and you will see the cost printed to the screen.
You should expect to see a cost of 32.07.

#### Gradient descent
Next, you will implement gradient descent in the file `gradientDescent.go`.
You need to supply the updates to `theta` within each iteration.

As you program, make sure you understand what you are trying to opti-
mize and what is being updated. Keep in mind that the cost `J(theta)` is parame-
terized by the vector `theta`, not `xys`. That is, we minimize the value of `J(theta)`
by changing the values of the vector `theta`, not by changing `xys`.

A good way to verify that gradient descent is working correctly is to look
at the value of `J(theta)` and check that it is decreasing with each step. The
starter code for gradientDescent.m calls computeCost on every iteration
and prints the cost. Assuming you have implemented gradient descent and
computeCost correctly, your value of `J(theta)` should never increase, and should
converge to a steady value by the end of the algorithm.

After you are finished, ex1.m will use your final parameters to plot the linear fit.

Your final values for `theta` will also be used to make predictions on profits in
areas of 35,000 and 70,000 people.