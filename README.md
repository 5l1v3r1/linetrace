# Goal

It is easy for a human to trace out lines when they look at a handwritten letter or digit. In general, we can easily see how to trace out a line drawing. But can a computer do it? The goal of this project is to train a computer to trace out lines in an image.

# Current solution

The current solution uses gradient descent on the coordinates of vertices in potential paths. It does not do incredibly well, but here are some of the better reproduction results I've gotten:

<table>
  <tr>
    <td>Original</td>
    <td>Reconstructed Path</td>
  </tr>
  <tr>
    <td><img src="results/0_orig.png"></td>
    <td><img src="results/0_repro.png"></td>
  </tr>
  <tr>
    <td><img src="results/1_orig.png"></td>
    <td><img src="results/1_repro.png"></td>
  </tr>
  <tr>
    <td><img src="results/2_orig.png"></td>
    <td><img src="results/2_repro.png"></td>
  </tr>
  <tr>
    <td><img src="results/3_orig.png"></td>
    <td><img src="results/3_repro.png"></td>
  </tr>
  <tr>
    <td><img src="results/4_orig.png"></td>
    <td><img src="results/4_repro.png"></td>
  </tr>
  <tr>
    <td><img src="results/5_orig.png"></td>
    <td><img src="results/5_repro.png"></td>
  </tr>
</table>
