1. Install Node Js if it's already not installed in the Ubuntu/Linux/mac.

2. change into the caliper-benchmarks directoy.

If you're in the test-network directory then perform

```bash
cd ../../
```


```bash
mkdir fabric-benchmark && cd fabric-benchmark
```

3. Clone caliper-benchmarks (ensure it is created inside the directory you created above)

```bash
git clone https://github.com/hyperledger/caliper-benchmarks
```

4. Change into the caliper-benchmarks directoy

```bash
cd caliper-benchmarks
```

5. Install the latest version of caliper

```bash
npm install --only=prod @hyperledger/caliper-cli
```

6. Bind to the fabric with the respective version

To bind with fabric 2.4 and utilise the new peer-gateway service introduced in fabric 2.4 run:

```bash
npx caliper bind --caliper-bind-sut fabric:2.4
```

If you wish to change the version of binding make sure to unbind your current binding (for example if you bound to fabric:2.2 unbind first with `caliper unbind --caliper-bind-sut fabric:2.2`) before binding to the new one.

7. Copy the required configuration files.

Run the command to know the container names which you're running.

```bash
docker ps -a --format "{{.Names}}" | grep -i dev-peer
```

Replace the output of the above command in `./caliper/networks/test-network.yaml` in the line numbers 21 and 22. Don't remove `- ` symbol.

You need to perform whenever the network has been down and then made it up. The container names will keep changing whenever you redeploy the network. 

Run the command to know the container names which you're running.

```bash
ls ../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/
```

'../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/<OUTPUT>'

Replace the output of the above command in `./caliper/networks/test-network.yaml` in the line number 35. Don't remove `- ` symbol.

Run the command to know the container names which you're running.

```bash
ls ../../fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore/
```

'../../fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore/<OUTPUT>'

Replace the output of the above command in `./caliper/networks/test-network.yaml` in the line number 46. Don't remove `- ` symbol.

8. Copy the network configuration file to your directory.

```bash
cd networks
```

Create a file `test-network.yaml` file and paste the content here. If the file already exists, then replace the content of the file.

9. Copy the benchmark files.

```bash
cd ../benchmarks
mkdir mitigation
cd mitigation
```

Copy all files in the `./caliper/benchmarks/` into the current directory. If the same files are existing then no need to do anything.

10. Benchmark execution

```bash
cd ../../
```

Ensure you are in the `caliper-benchmarks` directory before running the following command. You can use `pwd` to check the current working directory.

```bash
npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/test-network.yaml --caliper-benchconfig benchmarks/mitigation/config.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled
```


npx caliper launch manager:

Launches the Caliper benchmark manager, which is responsible for initializing and running the benchmark test.

--caliper-workspace ./:

Specifies the directory where Caliper's configuration files and workload modules are located. In this case, ./ points to the current directory.

--caliper-networkconfig test-network.yaml:

Points to the network configuration file. test-network.yaml contains the configuration for the Hyperledger Fabric network, such as peer organizations, channels, and other network details.

--caliper-benchconfig benchmark/config.yaml:

Points to the benchmark configuration file. benchmark/config.yaml defines the workload (benchmark tests) and other performance testing parameters (such as transaction frequency, number of rounds, etc.).

--caliper-flow-only-test:

Specifies that the benchmark should run the test only without any setup or tear down (i.e., it skips the initialization and smart contract installation phases).

--caliper-fabric-gateway-enabled:

Enables the use of the Fabric Gateway, which is the new way to interact with the Fabric network introduced in Fabric 2.0. This flag indicates that the tests should be conducted using the Fabric Gateway instead of the older Fabric SDK.

Breakdown of Additional Flags:
--caliper-report file:

Tells Caliper to generate a report and save it to a file.

--caliper-report-format json:

Specifies the format of the report. You can also use html, csv, or other formats supported by Caliper. In this case, we use JSON format.

--caliper-report-path ./test-results:

Specifies the directory where the report should be saved. In this case, the directory ./test-results will store the test results. You can change the path to any directory where you'd like to save the results.

11. This will generate `report.html` file.
12. To create visualization charts, use the `python` scripts in the `scripts` folder.